package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
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

type WebPlatformTest struct {
	url     string
	options options
}

func (s WebPlatformTest) Run(
	ctx context.Context,
	l *slog.Logger,
) (res []WebPlatformTestCase, err error) {
	defer func() {
		if s.options.ignorePanics {
			return
		}
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
				return
			}
			err = fmt.Errorf("panic running test: %s\n%v\n", s.url, r)
		}
	}()
	b := browser.New(
		browser.WithScriptEngine(v8engine.DefaultEngine()),
		browser.WithContext(ctx),
		browser.WithLogger(l),
	)
	defer b.Close()
	var win html.Window

	win, err = b.Open(s.url)
	l.Debug("Window opened")
	if err == nil {
		err = win.Clock().ProcessEvents(ctx)
	}
	l.Debug("Events processed")
	if err != nil {
		return
	}
	return s.parseResults(win)
}

func (s WebPlatformTest) parseResults(win html.Window) (res []WebPlatformTestCase, err error) {
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

func RunTestCase(
	ctx context.Context, tc TestCase, log *slog.Logger, o options,
) ([]WebPlatformTestCase, error) {
	path := tc.Path
	log = log.With(slog.String("TestCase", path))
	log.Info("Start test")

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	suite := WebPlatformTest{
		url:     tc.URL(),
		options: o,
	}
	return suite.Run(ctx, log)
}

type testResult struct {
	testCase TestCase
	res      []WebPlatformTestCase
	err      error
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
func testResults(
	ctx context.Context,
	tests <-chan TestCase,
	log *slog.Logger,
	o options,
) <-chan pendingTest {
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

				testCaseRes, err := RunTestCase(ctx, testCase, log, o)
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

type options struct {
	logger       *slog.Logger
	file         string
	includes     []string
	ignorePanics bool
}

func (o options) Logger() (res *slog.Logger) {
	if res = o.logger; res == nil {
		res = slog.Default()
	}
	return
}

type staticTestCaseSource struct {
	files    []string
	baseHref string
}

func (s staticTestCaseSource) testCases(ctx context.Context) <-chan TestCase {
	res := make(chan TestCase)
	go func() {
		defer close(res)
		for _, file := range s.files {
			res <- TestCase{
				PathElements: strings.Split(file, "/"),
				Path:         file,
			}
		}
	}()
	return res
}

func loadTestSource(o options) testCaseLoader {
	if o.file != "" {
		return staticTestCaseSource{
			baseHref: "https://wpt.live/",
			files:    []string{o.file},
		}
	} else {
		return manifestTestCaseSource{
			href:    "https://wpt.live/MANIFEST.json",
			options: o,
			filter: pathFilter{
				included: o.includes,
				excluded: []string{}}.isMatch,
		}
	}
}

// testCaseLoader is the interface for the method testCases() that return a
// channel of TestCases. The source should ideally return a buffered channel.
// The buffer size determines how many test cases can run in parallel
type testCaseLoader interface {
	testCases(ctx context.Context) <-chan TestCase
}

func parseOptions() options {
	var o = options{
		logger: newLogger(),
	}
	flag.StringVar(&o.file, "file", "", "")
	flag.Parse()
	o.includes = flag.Args()
	return o
}

func main() {
	var o = parseOptions()
	var source = loadTestSource(o)

	var logger = o.Logger()
	body := element("body")
	body.AppendChild(element("h1", "Web Application Test Report"))
	var errs []error
	var prevHeaders []string
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for pending := range testResults(ctx, source.testCases(ctx), logger, o) {
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
	// h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	// 	Level: slog.LevelInfo,
	// 	ReplaceAttr: log.ReplaceStackAttr,
	// })
	return slog.New(h)
}
