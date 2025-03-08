# Developer docs

This consists of a few separate documents

- [Code architecture](./Code-architecture.md) describes the overall dependency of packages.
- [Code structure](./Code-structure.md) - more detailed description of the packages.
- [DOM node implementation](./DOM-nodes.md) - The DOM tree itself (see below)
- [Testing](./Testing.md) - Guidelines for testing

## DOM Tree - Unidiomatic Go

Idiomatic Go uses small interfaces, and this leads to many elegant solutions.

But the problem domain itself, implementing browser behaviour **and implementing
the standards** does place a significant constraint on the types of solutions
applicable.

As a result, there are many types, and specifically interfaces, that are larger
than desired. This is particularly the case in the DOM tree itself.

## Polyfills

Where possible, JavaScript APIs are implemented through polyfills.

- Go code to install polyfills: https://github.com/gost-dom/browser/blob/main/scripting/v8host/polyfills.go 
- Folder for javascript files to load: https://github.com/gost-dom/browser/tree/main/scripting/v8host/polyfills
