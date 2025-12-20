package html

import (
	"fmt"

	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/url"
)

type location struct {
	*url.URL
}

func newLocation(u *url.URL) Location {
	return location{u}
}

func (l location) AncestorOrigins() DOMStringList {
	return nil
}

func (l location) Assign(string) error {
	return fmt.Errorf(
		"html/Location.Assign: Not implemented. %s",
		constants.MISSING_FEATURE_ISSUE_URL,
	)
}
func (l location) Replace(string) error {
	return fmt.Errorf(
		"html/Location.Replace: Not implemented. %s",
		constants.MISSING_FEATURE_ISSUE_URL,
	)
}

func (l location) Reload() error {
	return fmt.Errorf(
		"html/Location.Reload: Not implemented. %s",
		constants.MISSING_FEATURE_ISSUE_URL,
	)
}
