package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Marlliton/slogpretty"
	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/v8engine"
	xhtml "golang.org/x/net/html"
)

type WebPlatformTest string

func (s WebPlatformTest) Run(
	ctx context.Context,
	l *slog.Logger,
) (res []WebPlatformTestCase, err error) {
	b := browser.New(
		browser.WithScriptEngine(v8engine.DefaultEngine()),
		browser.WithContext(ctx),
		browser.WithLogger(l),
	)
	defer b.Close()
	var win html.Window

	win, err = b.Open(string(s))
	l.Debug("Window opened")
	if err == nil {
		err = win.Clock().ProcessEvents(ctx)
	}
	l.Debug("Events processed")
	if err != nil {
		return
	}
	r, _ := win.Document().QuerySelectorAll("#results > tbody > tr")
	rows := r.All()
	res = make([]WebPlatformTestCase, len(rows))
	var errs []error
	for i, row := range rows {
		tds, _ := row.(dom.Element).QuerySelectorAll("td")

		result := WebPlatformTestCase{
			Success: tds.Item(0).TextContent() == "Pass",
			Name:    tds.Item(1).TextContent(),
			Msg:     tds.Item(2).(dom.Element).OuterHTML(),
		}
		res[i] = result
		if !result.Success {
			errs = append(errs, fmt.Errorf("fail: %s", result.Name))
		}
	}
	err = errors.Join(errs...)
	return
}

type WebPlatformTestCase struct {
	Name    string
	Success bool
	Msg     string
}

func RunTestCase(tc TestCase, log *slog.Logger) ([]WebPlatformTestCase, error) {
	path := tc.Path
	log = log.With(slog.String("TestCase", path))
	log.Info("Start test")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	suite := WebPlatformTest(tc.URL())
	return suite.Run(ctx, log)
}

// includeList specifies the subpaths of the "testharness" tests that will be
// included.
var includeList = []string{
	"dom/nodes/Document-getElementById",
}

// excludeList filters individual subpaths that would have been included from
// the includeList.
var excludeList = []string{
	// "dom/abort/",
	// "dom/attributes-are-nodes",
}

type testResult struct {
	testCase TestCase
	res      []WebPlatformTestCase
	err      error
}

func filteredTests(ctx context.Context, r io.Reader, l *slog.Logger) <-chan TestCase {
	ch := make(chan TestCase, 8)
	go func() {
		defer func() { close(ch) }()
	testCaseLoop:
		for testCase := range ParseManifest(ctx, r, l) {
			path := testCase.Path
			for _, include := range includeList {
				if strings.HasPrefix(path, include) {
					for _, exclude := range excludeList {
						if strings.HasPrefix(path, exclude) {
							continue testCaseLoop
						}
					}

					ch <- testCase
					continue testCaseLoop
				}
			}

		}
	}()
	return ch
}

// pendingTest represent a test that was started. When the test has completed,
// the result can be read from the channel.
type pendingTest <-chan testResult

func (t pendingTest) result() testResult {
	return <-t
}

// testResults return a channel of pending test results. The return type is a
// channel of channels in order to have the channel reflect the order in which
// tests were started, not when they were completed.
func testResults(tests <-chan TestCase, log *slog.Logger) <-chan pendingTest {
	res := make(chan pendingTest, 64)
	go func() {
		var grp sync.WaitGroup
		defer func() { close(res) }()

		for testCase := range tests {
			grp.Add(1)
			resultChan := make(chan testResult)
			res <- resultChan
			go func() {
				defer grp.Add(-1)

				testCaseRes, err := RunTestCase(testCase, log)
				resultChan <- testResult{
					testCase: testCase,
					res:      testCaseRes,
					err:      err,
				}
			}()
		}
		grp.Wait()
	}()
	return res
}

func main() {
	logger := newLogger()
	logger.Info("Start")
	res, err := http.Get("https://wpt.live/MANIFEST.json")
	logger.Info("Got WPT manifest")
	if err != nil {
		panic(fmt.Sprintf("load manifest: %v", err))
	}
	defer res.Body.Close()

	body := element("body")
	body.AppendChild(element("h1", "Web Application Test Report"))

	var errs []error

	testCaseSource := filteredTests(context.Background(), res.Body, logger)
	var prevHeaders []string

	for pending := range testResults(testCaseSource, logger) {
		testCaseResult := pending.result()
		var (
			testCase = testCaseResult.testCase
			res      = testCaseResult.res
			err      = testCaseResult.err
		)

		for i, e := range testCase.PathElements {
			if len(prevHeaders) > i && prevHeaders[i] == e {
				if i != len(testCase.PathElements)-1 {
					continue
				}
			}
			prevHeaders = nil
			if i > 4 {
				i = 4 // There's no such thing as an H7
			}
			body.AppendChild(element(fmt.Sprintf("h%d", i+2), e))
		}
		prevHeaders = testCase.PathElements

		if err == nil {
			body.AppendChild(element("p", "All tests pass"))
			continue
		}

		body.AppendChild(
			element("div", element("span", "Full path: "), element("a",
				xhtml.Attribute{Key: "href", Val: testCase.URL()},
				xhtml.Attribute{Key: "target", Val: "_blank"},
				testCase.URL())),
		)

		logger.Error("ERROR", "err", err)
		errs = append(errs, err)

		tbl := element("table")
		thead := element("thead")
		tr := element("tr")
		tbody := element("tbody")
		body.AppendChild(tbl)
		tbl.AppendChild(thead)
		tbl.AppendChild(tbody)
		thead.AppendChild(tr)
		tr.AppendChild(element("th", "Success"))
		tr.AppendChild(element("th", "Test"))
		tr.AppendChild(element("th", "Msg"))

		for _, row := range res {
			tr := element("tr")
			pass := "FAIL"
			if row.Success {
				pass = "PASS"
			}
			tr.AppendChild(element("td", pass))
			tr.AppendChild(element("td", row.Name))
			if nodes, err := xhtml.ParseFragment(strings.NewReader(row.Msg), tr); err != nil {
				logger.Error("Error parsing node", "err", err, "fragment", row.Msg)
			} else {
				for _, node := range nodes {
					tr.AppendChild(node)
				}
			}

			tbody.AppendChild(tr)
		}
	}
	if err := save(body); err != nil {
		errs = append(errs, err)
	}
	if err := errors.Join(errs...); err != nil {
		os.Exit(1)
	}
}

func newLogger() *slog.Logger {
	opts := slogpretty.DefaultOptions()
	opts.Multiline = true
	opts.Level = slog.LevelInfo
	h := slogpretty.New(os.Stdout, opts)
	return slog.New(h)
}
