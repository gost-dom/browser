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
	"github.com/gost-dom/browser/scripting/sobekengine"
	xhtml "golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type WebPlatformTest string

func (s WebPlatformTest) Run(
	ctx context.Context,
	l *slog.Logger,
) (res []WebPlatformTestCase, err error) {
	b := browser.New(
		browser.WithScriptEngine(sobekengine.DefaultEngine()),
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

type TestCaseResult struct {
}

func RunTestCase(tc TestCase, log *slog.Logger) ([]WebPlatformTestCase, error) {
	path := tc.Path
	log = log.With(slog.String("TestCase", path))
	log.Info("Start test")

	url := fmt.Sprintf("https://wpt.live/%s", path)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	suite := WebPlatformTest(url)
	return suite.Run(ctx, log)
}

// includeList specifies the subpaths of the "testharness" tests that will be
// included.
var includeList = []string{
	"dom/nodes/Document-getElementById",
	// "fetch/api",
}

// excludeList filters individual subpaths that would have been included from
// the includeList.
var excludeList = []string{
	// "dom/abort/",
	// "dom/attributes-are-nodes",
}

type result struct {
	testCase TestCase
	res      []WebPlatformTestCase
	err      error
}

func filteredTests(ctx context.Context, r io.Reader) <-chan TestCase {
	ch := make(chan TestCase, 8)
	go func() {
		defer func() { close(ch) }()
	testCaseLoop:
		for testCase := range ParseManifest(ctx, r) {
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

func testResults(tests <-chan TestCase, log *slog.Logger) <-chan chan result {
	res := make(chan chan result, 64)
	go func() {
		var grp sync.WaitGroup
		defer func() { close(res) }()

		for testCase := range tests {
			grp.Add(1)
			resultChan := make(chan result)
			res <- resultChan
			go func() {
				defer grp.Add(-1)

				testCaseRes, err := RunTestCase(testCase, log)
				resultChan <- result{
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
	log := newLogger()
	res, err := http.Get("https://wpt.live/MANIFEST.json")
	if err != nil {
		panic(fmt.Sprintf("load manifest: %v", err))
	}
	defer res.Body.Close()

	body := element("body")
	body.AppendChild(element("h1", "Web Application Test Report"))

	var errs []error

	testCaseSource := filteredTests(context.Background(), res.Body)
	for testCaseResultCh := range testResults(testCaseSource, log) {
		testCaseResult := <-testCaseResultCh
		var (
			testCase = testCaseResult.testCase
			res      = testCaseResult.res
			err      = testCaseResult.err
		)
		body.AppendChild(element("h2", testCase.Path))
		if err == nil {
			body.AppendChild(element("p", "All tests pass"))
			continue
		}

		log.Error("ERROR", "err", err)
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
				log.Error("Error parsing node", "err", err, "fragment", row.Msg)
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

func save(body *xhtml.Node) (err error) {
	f, err := os.Create("report.html")
	if err == nil {
		doc := &xhtml.Node{Type: xhtml.DocumentNode}
		html := element("html")
		doc.AppendChild(html)
		html.AppendChild(body)
		err = xhtml.Render(f, doc)
	}
	return err
}

func newLogger() *slog.Logger {
	opts := slogpretty.DefaultOptions()
	opts.Multiline = true
	opts.Level = slog.LevelWarn
	h := slogpretty.New(os.Stdout, opts)
	return slog.New(h)
}

func element(tag string, opts ...any) *xhtml.Node {
	res := &xhtml.Node{
		Type:     xhtml.ElementNode,
		DataAtom: atom.Lookup([]byte(tag)),
		Data:     tag,
	}
	for _, opt := range opts {
		switch t := opt.(type) {
		case string:
			res.AppendChild(&xhtml.Node{
				Type: xhtml.TextNode,
				Data: t,
			})
		default:
			panic(fmt.Sprintf("invalid element option: %v", opt))
		}
	}
	return res
}
