package main_test

import (
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/testing/gosttest"
)

var tests = []string{
	"https://wpt.live/dom/nodes/Document-getElementById.html",
}

func TestWebApplicationTests(t *testing.T) {
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			t.Parallel()
			b := browser.New(
				browser.WithLogger(
					gosttest.NewTestingLogger(t), // gosttest.MinLogLevel(slog.LevelDebug),

				),
			)
			t.Cleanup(b.Close)
			win, err := b.Open("https://wpt.live/dom/nodes/Document-getElementById.html")
			if err != nil {
				t.Fatalf("Error opening window: %v", err)
			}
			if err = win.Clock().ProcessEvents(t.Context()); err != nil {
				t.Fatalf("Error processing events: %v", err)
			}

			rows, _ := win.Document().QuerySelectorAll("#results > tbody > tr")
			for _, row := range rows.All() {
				tds, _ := row.(dom.Element).QuerySelectorAll("td")
				result := tds.Item(0).TextContent()
				name := tds.Item(1).TextContent()
				if result == "Fail" {
					errMsg := tds.Item(2).TextContent()
					t.Errorf("Fail: %s\n%s", name, errMsg)
				}
			}
		})
	}
}
