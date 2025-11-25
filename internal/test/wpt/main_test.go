package main_test

import (
	"testing"

	wpt "github.com/gost-dom/browser/internal/test/wpt"
	"github.com/gost-dom/browser/testing/gosttest"
)

var tests = []string{
	"https://wpt.live/dom/nodes/Document-getElementById.html",
}

func TestWebApplicationTests(t *testing.T) {
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			t.Parallel()

			l := gosttest.NewTestingLogger(t)
			wptSuite := wpt.WebPlatformTest(test)
			cases, err := wptSuite.Run(t.Context(), l)
			if err != nil {
				t.Fatalf("Error running: %v\n", err)
			}
			for _, c := range cases {
				if !c.Success {
					t.Errorf("%s: %s\n", c.Name, c.Msg)
				}
			}
		})
	}
}
