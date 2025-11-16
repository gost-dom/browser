package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gost-dom/browser"
)

func main() {
	h := slog.NewTextHandler(os.Stdout, nil)
	log := slog.New(h)
	b := browser.New(browser.WithLogger(log))
	ctx := context.Background()
	win, err := b.Open("https://wpt.live/dom/nodes/CharacterData-appendChild.html")
	if err != nil {
		log.Error("Error opening window", "err", err)
		os.Exit(1)
	}
	err = win.Clock().ProcessEvents(ctx)
	if err != nil {
		log.Error("Error processing events", "err", err)
	}
	log.Info("Document", "html", win.Document().DocumentElement().OuterHTML())
}
