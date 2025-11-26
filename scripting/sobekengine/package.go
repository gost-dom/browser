// -- go:generate ../../internal/code-gen/code-gen -g goja

// The sobekhost package provides functionality to execute client-scripts in gost-dom.
//
// This package uses [sobek] as the script engine, a pure Go JavaScript engine.
//
// If some JavaScript features are not working as intended, [v8host] proves a
// script host using the v8 engine; which will support any valid (browser
// supported) JavaScript that you can throw at it; but has significant overhead,
// and relies on CGo.
//
// See also: https://github.com/grafana/sobek
package sobekengine
