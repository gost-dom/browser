package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"iter"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/Marlliton/slogpretty"
	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/sobekengine"
)

//go:embed manifest.json
var manifest []byte

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
	for i, row := range rows {
		tds, _ := row.(dom.Element).QuerySelectorAll("td")
		res[i] = WebPlatformTestCase{
			Success: tds.Item(0).TextContent() == "Pass",
			Name:    tds.Item(1).TextContent(),
			Msg:     tds.Item(2).(dom.Element).OuterHTML(),
		}
	}
	return
}

type WebPlatformTestCase struct {
	Name    string
	Success bool
	Msg     string
}

func LoadManifest() (m Manifest, err error) {
	err = json.Unmarshal(manifest, &m)
	return
}

func RunTestCase(tc TestCase, log *slog.Logger) error {
	path := tc.Path
	log = log.With(slog.String("TestCase", path))
	log.Info("Start test")

	url := fmt.Sprintf("https://wpt.live/%s", path)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	suite := WebPlatformTest(url)
	res, err := suite.Run(ctx, log)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}

	var errs []error
	for _, row := range res {
		if row.Success {
			log.Info("PASS", "test", row.Name)
		} else {
			log.Error("FAIL", "test", row.Name)
			errs = append(errs, fmt.Errorf("FAIL: %s: %s", path, row.Name))
		}
	}
	return errors.Join(errs...)
}

var includeList = []string{
	"dom/nodes/Document-getElementById",
}

var excludeList = []string{
	// "dom/abort/",
	// "dom/attributes-are-nodes",
}

func filteredTests(m Manifest) iter.Seq[TestCase] {
	return func(yield func(TestCase) bool) {
	testCaseLoop:
		for testCase := range m.All() {
			path := testCase.Path
			for _, include := range includeList {
				if strings.HasPrefix(path, include) {
					for _, exclude := range excludeList {
						if strings.HasPrefix(path, exclude) {
							continue testCaseLoop
						}
					}
					if !yield(testCase) {
						return
					}
					continue testCaseLoop
				}
			}
			// if strings.HasPrefix(testCase.Path, "dom/") {
			// }
		}
	}
}

func main() {
	opts := slogpretty.DefaultOptions()
	opts.Multiline = true
	// opts.Level = slog.LevelDebug
	h := slogpretty.New(os.Stdout, opts)
	log := slog.New(h)

	manifest, err := LoadManifest()
	if err != nil {
		log.Error("Error", "err", err)
	}
	var errs []error

	for testCase := range filteredTests(manifest) {
		if err := RunTestCase(testCase, log); err != nil {
			// For now, bail early.
			log.Error("ERROR", "err", err)
			os.Exit(1)
			errs = append(errs, err)
		}
	}
	if err := errors.Join(errs...); err != nil {
		os.Exit(1)
	}
}
