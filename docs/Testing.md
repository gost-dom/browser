# Testing

All code must be covered by tests

But the package is a mix of different paradigms, so what to use?

**Don't use Ginkgo**

Existing Ginkgo tests will eventually be migrated to non-Ginko tests. Currently
the replacement is largely based on [Testify
suites](https://pkg.go.dev/github.com/stretchr/testify/suite) that integrate
with Go's subtest functionality. Contributor feedback will affect the future of
this.

Gingko was used historically, but this was a mistake, and it's a barrier for new
contributors to join.[^1]

## Assertions

You are free to write assertions in any way you like.

Gomega has been used historically, and provides expressive assertions through
custom matchers, but the true benefit isn't achieved until you take your time to
write the custom matchers.

## Testing stripped HTML documents

Many tests create a `Document` by parsing an HTML string. The HTML specification
dictates implicitly creating `<html>`, `<head>`, and `<body>` if these elements
are missing in the source.

For brevity, `<html>`, and `<head>` are left out where they don't affect the
outcome of the test, but `<body>` is kept to explicitly state that the string
contains the _entire_ body of the test document.

Tests can safely rely on this behaviour, as it is according to standards.

---

[^1]: When Ginkgo was first conceived, it provided some benefits for organising
    larget test suites. But these have diminished with advancements in Go's own
    testing capabilities.


