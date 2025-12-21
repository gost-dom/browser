package html

import (
	"fmt"

	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/url"
)

type location struct {
	entity.Entity
	*url.URL
}

func urlToLocation(u *url.URL) *location {
	if u == nil {
		return nil
	}
	return &location{URL: u}
}

func (l *location) set(u *url.URL) { l.URL = u }

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
