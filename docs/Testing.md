# Testing

All code must be covered by tests!

Some tests use testify suites, but native Go tests are preferred.
Nested/table-driven tests are encouraged where it makes sense.

## Assertions

Feel free to write assertions in any way you like; but avoid unnecessary
duplication. 

Gomega has been used historically, and provides expressive assertions through
custom matchers, but the true benefit isn't achieved until you take your time to
write the custom matchers.

Gost-DOM includes some custom matchers useful for verifying properties of the
DOM.

## Testing stripped HTML documents

Many tests create a `Document` by parsing an HTML string. The HTML specification
dictates implicitly creating `<html>`, `<head>`, and `<body>` if these elements
are missing in the source.

For brevity, `<html>`, and `<head>` are left out where they don't affect the
outcome of the test, but `<body>` is kept to explicitly state that the string
contains the _entire_ body of the test document.

Tests can safely rely on this behaviour, as this is specified by whatwg
living documentation.
