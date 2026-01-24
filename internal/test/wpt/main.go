package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"strconv"
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

// Start more concurrent tests than we have CPUs. Each test has idle time, e.g.,
// waiting for HTTP responses when downloading the individual test case.
var MAX_CONCURRENT_TESTS = runtime.NumCPU()

type WptSuiteResult struct {
	TestCases   []WebPlatformTestCase
	TotalCount  int
	PassedCount int
	Error       error
}

type WebPlatformTest struct {
	url     string
	options options
}

func (s WebPlatformTest) Run(
	ctx context.Context,
	l *slog.Logger,
) (res WptSuiteResult) {
	defer func() {
		if s.options.ignorePanics {
			return
		}
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				res.Error = e
				return
			}
			res.Error = fmt.Errorf("panic running test: %s\n%v\n", s.url, r)
		}
	}()
	b := browser.New(
		browser.WithScriptEngine(v8engine.DefaultEngine()),
		browser.WithContext(ctx),
		browser.WithLogger(l),
	)
	defer func() {
		defer func() {
			if err := recover(); err != nil {
				l.Error("Error closing browser", slog.Any("err", err))
			}
		}()
		b.Close()
	}()

	var win html.Window

	l.Debug("Open window", slog.String("url", s.url))
	win, res.Error = s.openWindow(b)
	if res.Error == nil {
		l.Debug("ProcessEvents")
		res.Error = s.processEvents(ctx, win)
	}
	if res.Error != nil {
		return
	}
	return s.parseResults(win)
}

func (t WebPlatformTest) openWindow(b *browser.Browser) (html.Window, error) {
	win, err := b.Open(t.url)
	if err != nil {
		err = fmt.Errorf("Open window: %w", err)
	}
	return win, err
}

func (t WebPlatformTest) processEvents(ctx context.Context, win html.Window) error {
	err := win.Clock().ProcessEvents(ctx)
	if err != nil {
		err = fmt.Errorf("Process events: %w", err)
	}
	return err
}

func (s WebPlatformTest) parseResults(win html.Window) (res WptSuiteResult) {
	r, _ := win.Document().QuerySelectorAll("#results > tbody > tr")
	rows := r.All()
	res = WptSuiteResult{
		TestCases: make([]WebPlatformTestCase, len(rows)),
	}
	var errs []error
	for i, row := range rows {
		tds, _ := row.(dom.Element).QuerySelectorAll("td")

		result := WebPlatformTestCase{
			Success: tds.Item(0).TextContent() == "Pass",
			Name:    tds.Item(1).TextContent(),
			Msg:     tds.Item(2).(dom.Element).OuterHTML(),
		}
		res.TestCases[i] = result

		res.TotalCount++
		if result.Success {
			res.PassedCount++
		} else {
			errs = append(errs, fmt.Errorf("fail: %s", result.Name))
		}
	}

	res.Error = errors.Join(errs...)
	return
}

type WebPlatformTestCase struct {
	Name    string
	Success bool
	Msg     string
}

func RunTestCase(
	ctx context.Context, tc TestCase, log *slog.Logger, o options,
) (res WptSuiteResult, err error) {
	log.Info("Start test")
	ch := make(chan WptSuiteResult, 1)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	go func() {
		suite := WebPlatformTest{
			url:     tc.URL(),
			options: o,
		}
		select {
		case ch <- suite.Run(ctx, log):
		case <-ctx.Done():
		}
	}()
	select {
	case res := <-ch:
		return res, nil
	case <-ctx.Done():
		err = errors.New("Context aborted")
		return
	}
}

type testResult struct {
	testCase TestCase
	res      WptSuiteResult
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
	res := make(chan pendingTest, MAX_CONCURRENT_TESTS)
	go func() {
		var grp sync.WaitGroup
		defer func() { close(res) }()

		for testCase := range tests {
			log := log.With(slog.String("TestCase", strings.TrimSpace(testCase.Path)))
			resultChan := make(chan testResult)
			select {
			case res <- resultChan:
				grp.Go(func() {

					testCaseRes, err := RunTestCase(ctx, testCase, log, o)
					if err == nil && len(testCaseRes.TestCases) == 0 {
						err = errors.New("Test suite returned no results")
					}
					if err == nil {
						err = testCaseRes.Error
					}
					resultChan <- testResult{
						testCase: testCase,
						res:      testCaseRes,
						err:      err,
					}
				})
			case <-ctx.Done():
				return
			}
		}
		grp.Wait()
	}()
	return res
}

type options struct {
	logger       *slog.Logger
	logLevel     string
	logType      string
	logNoColor   bool
	file         string
	includes     []string
	ignorePanics bool
}

func (o *options) initLogger() {
	var h slog.Handler
	if o.logType == "pretty" {
		opts := slogpretty.DefaultOptions()
		opts.Multiline = true
		opts.Level = o.LogLevel()
		opts.Colorful = !o.logNoColor
		h = slogpretty.New(os.Stdout, opts)
	}
	if h == nil {
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:       o.LogLevel(),
			ReplaceAttr: removeTime,
		})
	}
	o.logger = slog.New(h)
}

func removeTime(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	return a
}

func (o options) LogLevel() slog.Level {
	switch strings.ToLower(o.logLevel) {
	case "debug", "d":
		return slog.LevelDebug
	case "info", "i":
		return slog.LevelInfo
	case "warn", "w":
		return slog.LevelWarn
	case "error", "err", "e":
		return slog.LevelError
	}
	if lvl, err := strconv.Atoi(o.logLevel); err == nil {
		return slog.Level(lvl)
	}
	return slog.LevelWarn
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

func (s staticTestCaseSource) testCases(
	ctx context.Context,
	cancel context.CancelCauseFunc,
) <-chan TestCase {
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
				excluded: excludedSuites,
			}.isMatch,
		}
	}
}

// testCaseLoader is the interface for the method testCases() that return a
// channel of TestCases. The source should ideally return a buffered channel.
// The buffer size determines how many test cases can run in parallel.
//
// If an error occurrs during loading tests cases, e.g., a HTTP disconnection,
// or invalid JSON is returned, the implementer must close the returned channel
// and call cancelCause with an error describing the calse
type testCaseLoader interface {
	testCases(ctx context.Context, cancelCause context.CancelCauseFunc) <-chan TestCase
}

func parseOptions() options {
	var o options
	flag.StringVar(&o.file, "file", "", "")
	flag.StringVar(&o.logLevel, "log-level", "warn", "")
	flag.StringVar(&o.logType, "log-type", "", "")
	flag.BoolVar(&o.logNoColor, "no-color", false, "")
	flag.Parse()
	o.initLogger()
	o.includes = flag.Args()
	return o
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	ctx, cancelCause := context.WithCancelCause(ctx)
	defer stop()

	var o = parseOptions()
	var source = loadTestSource(o)

	var logger = o.Logger()
	body := element("body")
	summary := element("div")
	body.AppendChild(element("h1", "Web Application Test Report"))
	body.AppendChild(summary)
	var errs []error
	var prevHeaders []string

	suiteCount := 0
	passedSuiteCount := 0
	testCount := 0
	passedTestCount := 0

	testCases := source.testCases(ctx, cancelCause)
	for pending := range testResults(ctx, testCases, logger, o) {
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

		suiteCount++
		testCount = testCount + res.TotalCount
		passedTestCount = passedTestCount + res.PassedCount
		if err == nil {
			body.AppendChild(element("p", "All tests pass"))
			passedSuiteCount++
			continue
		}
		logger.Error("ERROR", "err", err)
		errs = append(errs, err)

		body.AppendChild(
			element("div", element("span", "Full path: "), element("a",
				xhtml.Attribute{Key: "href", Val: testCase.URL()},
				xhtml.Attribute{Key: "target", Val: "_blank"},
				testCase.URL())),
		)

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

		for _, row := range res.TestCases {
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
	select {
	case <-ctx.Done():
		err := context.Cause(ctx)
		if err != nil && !errors.Is(err, context.Canceled) {
			errs = append(errs, err)
		}
	default:
	}
	summary.AppendChild(
		element("p", fmt.Sprintf("Passed test suites %d/%d", passedSuiteCount, suiteCount)),
	)
	summary.AppendChild(
		element("p", fmt.Sprintf("Passed test cases %d/%d", passedTestCount, testCount)),
	)
	if err := save(body); err != nil {
		errs = append(errs, err)
	}
	if err := errors.Join(errs...); err != nil {
		os.Exit(1)
	}
}
