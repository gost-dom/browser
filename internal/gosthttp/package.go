// Package gosthttp provides functionality to bypass the TCP stack.
//
// This includes the [TestRoundTripper] that provides an implementation od the
// http.RoundTripper interface which communicates directly with an http.Server
// instance.
package gosthttp
