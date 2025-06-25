// Package gomegamatchers just exposes gomega matchers for easier importing.
//
// Dot-importing the gomega package brings in all exported functions; which is
// undesirable when you just want the matchers.
package gomegamatchers

import "github.com/onsi/gomega"

var Equal = gomega.Equal
var HaveValue = gomega.HaveValue
var BeEquivalentTo = gomega.BeEquivalentTo
var BeTrue = gomega.BeTrue
var BeFalse = gomega.BeFalse

var ContainSubstring = gomega.ContainSubstring

var Succeed = gomega.Succeed
var HaveOccurred = gomega.HaveOccurred

var BeNil = gomega.BeNil
var BeEmpty = gomega.BeEmpty

var WithTransform = gomega.WithTransform

var ContainElement = gomega.ContainElement
var HaveKeyWithValue = gomega.HaveKeyWithValue
var ConsistOf = gomega.ConsistOf
var HaveExactElements = gomega.HaveExactElements
var And = gomega.And
var HaveField = gomega.HaveField
var ContainElements = gomega.ContainElements

var BeClosed = gomega.BeClosed
