// Package gomegamatchers just exposes gomega matchers for easier importing.
//
// Dot-importing the gomega package brings in all exported functions; which is
// undesirable when you just want the matchers.
package gomegamatchers

import "github.com/onsi/gomega"

var Equal = gomega.Equal
var BeEquivalentTo = gomega.BeEquivalentTo

var ContainSubstring = gomega.ContainSubstring

var Succeed = gomega.Succeed
var HaveOccurred = gomega.HaveOccurred
var MatchRegexp = gomega.MatchRegexp
