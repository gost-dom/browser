// Package interfaces is an internal package containing generated go interfaces
//
// These exist as a "guarantee" that types implementing a specific IDL interface
// has the correct methods. Some of these may be used in actual code, e.g.,
// javascript wrapper classes; and some may not even be used in real code; but
// with a test in this package to ensure that a type conforms to the interface.
package interfaces
