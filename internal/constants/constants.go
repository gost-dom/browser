// Package constants is a collection of values that are used many times in the
// implementation, but has no relevance to users of the library, e.g., a
// link to where you can file an issue when you encounter a not-implemented
// feature; or a feature that is not fully implemented, e.g. Element.matches
// that only support a subset of CSS selectors.
package constants

import "errors"

const PACKAGE_NAME = "Gost-DOM"

const MISSING_FEATURE_ISSUE_URL = "Create feature request here: https://github.com/gost-dom/browser/issues"

const BUG_ISSUE_URL = "This is a bug in " + PACKAGE_NAME + ". Please fill a bug report here: https://github.com/gost-dom/browser/issues"

var ErrGostDomBug = errors.New(BUG_ISSUE_URL)
var ErrGostDomMissingFeature = errors.New(MISSING_FEATURE_ISSUE_URL)

const BUG_ISSUE_DETAILS = "Please include the contents below, as well as other contextual information you might have"
