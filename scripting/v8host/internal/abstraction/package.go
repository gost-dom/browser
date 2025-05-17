// Package abstraction provides an abstraction on top of V8.
//
// The purpose is to allow a modular approach to implementing client-side APIs,
// where the modules don't need to address a specific script engine; and client
// code can witch between script engines.
//
// The ability to switch script engine seems like a really good idea, as you
// have two different priorities you might want to switch between:
//
// - Using an engine that you know contains lates ES standards (V8)
// - Using a pure Go engine, simplifying the build pipeline.
//
// Additionally, if other projects would support SpiderMonkey (Firefox) and
// JavaScriptCode (Safari), you could verify your code using the script engines
// of the different browsers. (Do note, that it is outside the scope of this
// project to simulate the differences of different browsers).
package abstraction
