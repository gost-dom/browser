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

## Commit guidelines

Change logs are generated automatically from git commit messages. [conventional
commit](https://www.conventionalcommits.org/en/v1.0.0/) defines the format for
commit messages.

Conventional commits need not be followed rigurously, but notes that should
appear in changelog **must** follow the format.[^1]

### Commit format

```
<type>(<optional scope>): Title

Description

<FOOTER>: Footer title

Footer description
```

- Commit title must be short.
- The optional exclamation mark in 
- If the title isn't enough, the body should describe _why_ the commit is good.
- Both title and body shall be written in the imperative format.

### Commit types

The commit types are based on the [Angular commit message
format](https://github.com/angular/angular/blob/main/contributing-docs/commit-message-guidelines.md#commit-message-format) with some changes

| Type         | Description                                                           |
|--------------|-----------------------------------------------------------------------|
| **build**    | Changes that affect the build scripts                                 |
| **ci**       | Changes to the CI configuration                                       |
| **cleanup**  | Remove dead code, print statements, or comments                       |
| **docs**     | Documentation only changes                                            |
| **drop**     | An existing feature is removed or no longer supported                 |
| **feat**     | A new feature                                                         |
| **fix**      | A bug fix                                                             |
| **perf**     | A code change that improves performance                               |
| **refactor** | A code change that improves code and **doesn't change behaviour**[^2] |
| **test**     | Adding missing tests, correcting existing tests, or refactors tests   |
| **work**     | An incremental change working towards a feature                       |

A commit may naturally fall into multiple categories. A feature must have tests,
and ideally, test and implementation are in the same commit of type `feat` or
`work`.

The `work` type is intended for incremental changes for a larger feature; but
the commits are on a different level than what we want to describe in the change
log.[^3]

### Commit scopes

| Scope       | Description                                           |
|-------------|-------------------------------------------------------|
| **dom**     | Internal DOM model                                    |
| **html**    | Internal HTML representation                          |
| **v8**      | V8 script host                                        |
| **readme**  | Readme file changes. Only applicable for `docs` type  |
| **codegen** | Changes to the code generator                         |

A "feature" scoped with `dom` or `html` signifies that any relevant JavaScript
bindings are missing.

### High quality PR

The guidelines for a larger feature is

- create multiple `work` commits with short messages
- refactor before feature
- rebase before merge
- one `feat` non-fast forward merge commit with detailed description of the feature.

Instead of _creating_ the `feat` merge commit, you can write the commit message in the
PR title/body.

By placing refactors _before_ the feature itself, they can be merged
individually, separating refactor commits from changes to behaviour.

```
* feat: Add feature Y (PR)
|\  
| * work: Add 'A' to API
| * work: Add 'B' to API
|/  
* refactor: Make A easier to implement
* refactor: Make B easier to implement
* feat: Add feature X
```

The diff of the feature commit itself will list the cumulative changes of all
the work commits.

---

[^1]: Changes that **must** generate changelog messages include new features, bug
    fixes, deprecations, breaking changes (API change or dropped features, or
    changed behaviour)

[^2]: This deviates from the angular convention, which defines refactor to mean
    something different that the commonly accepted meaning.

[^3]: Goja support is ongoing. It's not worth mentioning any incremental updates
    to Goja in the changelog until it is actually ready to use.
