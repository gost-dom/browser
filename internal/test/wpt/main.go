package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Marlliton/slogpretty"
	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/dom"
)

func main() {
	opts := slogpretty.DefaultOptions()
	opts.Level = slog.LevelDebug
	opts.Multiline = true
	h := slogpretty.New(os.Stdout, opts)
	log := slog.New(h)

	log.Info("Start test run")
	b := browser.New(browser.WithLogger(log))
	ctx := context.Background()
	win, err := b.Open("https://wpt.live/dom/nodes/Document-getElementById.html")

	if err != nil {
		log.Error("Error opening window", "err", err)
		os.Exit(1)
	}
	err = win.Clock().ProcessEvents(ctx)
	if err != nil {
		log.Error("Error processing events", "err", err)
	}

	rows, _ := win.Document().QuerySelectorAll("#results > tbody > tr")
	for _, row := range rows.All() {
		tds, _ := row.(dom.Element).QuerySelectorAll("td")
		result := tds.Item(0).TextContent()
		name := tds.Item(1).TextContent()
		if result == "Pass" {
			log.Info("PASS", "status", result, "test", name)
		} else {
			log.Error("FAIL", "status", result, "test", name)
		}
	}
}
