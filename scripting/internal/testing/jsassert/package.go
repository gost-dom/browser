// Package jsassert is a failed attempt to write assertions in JavaScript
//
// The package brings assertion functionality into JavaScript scope; and would
// tie them back to Go's testing.TB.
//
// Clever in itself; but the wrong failing source code line would show up in the
// test output, despite having called t.Helper() where possible.
package jsassert
