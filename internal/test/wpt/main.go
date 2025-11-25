package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/Marlliton/slogpretty"
	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/v8browser"
)

type WebPlatformTest string

func (s WebPlatformTest) Run(
	ctx context.Context,
	l *slog.Logger,
) (res []WebPlatformTestCase, err error) {
	b := v8browser.New(
		browser.WithContext(ctx),
		browser.WithLogger(l),
	)
	defer b.Close()
	var win html.Window

	win, err = b.Open("https://wpt.live/dom/nodes/Document-getElementById.html")
	if err == nil {
		err = win.Clock().ProcessEvents(ctx)
	}
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

func main() {
	opts := slogpretty.DefaultOptions()
	opts.Multiline = true
	h := slogpretty.New(os.Stdout, opts)
	log := slog.New(h)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	suite := WebPlatformTest("https://wpt.live/dom/nodes/Document-getElementById.html")
	res, err := suite.Run(ctx, log)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Test error: %v\n", err)
		os.Exit(1)
	}
	for _, row := range res {
		if row.Success {
			log.Info("PASS", "test", row.Name)
		} else {
			log.Error("FAIL", "test", row.Name)
		}
	}
}
