package html

import (
	"fmt"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/url"
)

// HistoryEventHashChange is fired at the window when the URL fragment changes as
// a same-document navigation (e.g. `location.hash = ...`).
const HistoryEventHashChange = "hashchange"

// HashChangeEventInit carries the fully-qualified URLs before and after the
// fragment change, matching the HashChangeEvent interface.
//
// See also: https://developer.mozilla.org/en-US/docs/Web/API/HashChangeEvent
type HashChangeEventInit struct {
	OldURL string
	NewURL string
}

type location struct {
	entity.Entity
	*url.URL
	// win is the owning window, set when the document adopts this location, so a
	// same-document navigation (SetHash) can dispatch a hashchange event. It may
	// be nil for a detached location, in which case SetHash mutates the URL
	// silently (as before).
	win *window
}

func urlToLocation(u *url.URL) *location {
	if u == nil {
		return nil
	}
	return &location{URL: u}
}

func (l *location) set(u *url.URL) { l.URL = u }

// SetHash implements the `location.hash` setter as a same-document navigation:
// it updates the fragment and, if the fragment actually changed, fires a
// hashchange event at the owning window. This overrides the SetHash promoted
// from the embedded *url.URL (which only mutates the URL).
func (l *location) SetHash(val string) {
	oldURL := l.URL.Href()
	oldHash := l.URL.Hash()
	l.URL.SetHash(val)
	if l.URL.Hash() == oldHash {
		return // no change → no event (per spec)
	}
	if l.win != nil {
		l.win.DispatchEvent(&event.Event{
			Type: HistoryEventHashChange,
			Data: HashChangeEventInit{OldURL: oldURL, NewURL: l.URL.Href()},
		})
	}
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
