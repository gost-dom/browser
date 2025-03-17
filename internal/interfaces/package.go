// Package interfaces contains go interfaces generated from IDL specs
//
// The generated interfaces guarantee that go interfaces correspond to IDL
// specifications. Not all interfaces are used in production code; but then they
// are used in test code, where a type assertion is used to guarantee that a go
// type corresponds to a concrete interface.
//
// This serves as a verification that web APIs are implemented according to the
// interfaces, as a test will fail if not; and can detect changes to the web IDL
// standards.
//
// The types are also kept in a separate package for that exact reason, to avoid
// polluting run-time packages with code that has not use at run-time.
package interfaces
