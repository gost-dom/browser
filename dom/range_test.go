package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/stretchr/testify/require"
)

func TestRange(t *testing.T) {
	doc := ParseHtmlString(`
		<body>
			<h1>Hello, World!</h1>
			<p>Testing ranges</p>
		</body>
	`)
	r := &dom.Range{}
	start, err1 := doc.QuerySelector("h1")
	end, err2 := doc.QuerySelector("p")

	require.NoError(t, err1)
	require.NoError(t, err2)
	require.NoError(t, r.SetStart(start, 0))
	require.NoError(t, r.SetEnd(end, 0))
}
