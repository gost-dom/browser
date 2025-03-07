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

---

[^1]: When Ginkgo was first conceived, provided some benefits. But these have
    diminished with advancements in Go's own testing capabilities.


