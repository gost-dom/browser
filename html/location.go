package html

import (
	"fmt"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/constants"

	netURL "net/url"
)

type location struct {
	dom.URL
	neturl *netURL.URL
}

func newLocation(u *netURL.URL) Location {
	return location{dom.NewURLFromNetURL(u), u}
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
