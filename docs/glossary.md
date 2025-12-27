# Glossary

- **Interface**. 

  The IDL Specification defines _interfaces_; which are exposed in certain
  scopes, implemented by "classes" in JavaScript.
  - The interfaces can be composed of _partial_ or _mixin_ interfaces.
  - IDL Interfaces and mixin interfaces are represented in Go, and typically
    exposed as Go `interface` types.

## JavaScript

- **Realm**
  
  An isolated JavaScript environment with it's own set of objects in global
  scope. When navigating, the browser creates a new realm.

  Note: I wasn't aware of this term from the start of the project, so it isn't
  used consistently where applicable.

  See alse: [tc39 definition of realm]

[tc39 definition of realm]: https://tc39.es/ecma262/multipage/executable-code-and-execution-contexts.html#realm
