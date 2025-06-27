// Package promise provides a Go way to represent the concept of a Promise.
//
// In JavaScript, a promise represents an operation that will fulfill with a
// value in the future, or produce an error. The outcome of the operation is
// represented by generic type [Result], containing either an error or a
// fulfilled value. The generic [Promise] type is just a channel where the value
// will be received.
//
// This package is not intended to provide be a full implementation of
// JavaScript promises. The intention is to simplify implementation of web APIs
// that returns promises, e.g. the [fetch API].
//
// [fetch API]: https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API
package promise
