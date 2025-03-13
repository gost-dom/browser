# Developer docs

This consists of a few separate documents

- [Code architecture](./Code-architecture.md) describes the overall dependency of packages.
- [Code structure](./Code-structure.md) - more detailed description of the packages.
- [DOM node implementation](./DOM-nodes.md) - The DOM tree itself (see below)
- [Testing](./Testing.md) - Guidelines for testing

## Commit guidelines

- Commit title must be short.
- If the title isn't enough, the body should describe _why_ the commit is good.
- Commit messages follow [conventional
commit](https://www.conventionalcommits.org/en/v1.0.0/)

Conventional commits are not followed rigurously, but change logs are based on
them, so the important tags are `feat:` and `fix:`.

**NOTE:** A high-quality PR would consist of clear individual commits describing
technical changes. The PR title and message can describe the feature, turning
the PR into the conventional commit for the changelog. 

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
