// Package cache provides the ability to cache expensive resources
//
// The cache is intended to not only cache HTTP response; but also expensive
// processing performed on the response. E.g., parsing JavaScript into a
// compiled AST.
package cache
