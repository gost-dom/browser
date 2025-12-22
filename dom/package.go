//go:generate ../internal/code-gen/code-gen -g gotypes -p dom

// Package dom provides the fundamental DOM implementation for Gost-DOM.
//
// The DOM includes a [Node] implementation, and it's fundamental special types,
// such as [Element], [Document], [Text], etc; as well as events througn the
// [EventTarget] type.
//
// Specific implementation of HTML element types, including the HTMLDocument, is
// in the html package.
package dom
