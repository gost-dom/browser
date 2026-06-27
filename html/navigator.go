package html

import "github.com/gost-dom/browser/internal/entity"

// NavigatorProfile describes the values reported by the window's navigator. It
// lets client code dictate the browser identity when creating a browser, rather
// than the core module hardcoding a single profile.
//
// Configure it with the browser package's WithNavigator option. Fields left as
// their zero value fall back to the corresponding [DefaultNavigatorProfile]
// value, so a partial profile only needs to set the fields it cares about.
type NavigatorProfile struct {
	UserAgent string
	Platform  string
	Vendor    string
	// Language defaults to the first entry of Languages when left empty.
	Language            string
	Languages           []string
	HardwareConcurrency int
}

// DefaultNavigatorProfile is the profile reported when none is configured. It
// describes a current Chrome release, so applications under test observe a
// realistic, modern browser. Override it with the browser package's
// WithNavigator option.
var DefaultNavigatorProfile = NavigatorProfile{
	UserAgent:           "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
	Platform:            "Linux x86_64",
	Vendor:              "Google Inc.",
	Language:            "en-US",
	Languages:           []string{"en-US", "en"},
	HardwareConcurrency: 8,
}

// withDefaults returns the profile with every zero-valued field replaced by the
// corresponding [DefaultNavigatorProfile] value, so a partial profile remains
// usable. An empty Language defaults to the first configured language.
func (p NavigatorProfile) withDefaults() NavigatorProfile {
	d := DefaultNavigatorProfile
	if p.UserAgent == "" {
		p.UserAgent = d.UserAgent
	}
	if p.Platform == "" {
		p.Platform = d.Platform
	}
	if p.Vendor == "" {
		p.Vendor = d.Vendor
	}
	if len(p.Languages) == 0 {
		p.Languages = d.Languages
	}
	if p.Language == "" {
		if len(p.Languages) > 0 {
			p.Language = p.Languages[0]
		} else {
			p.Language = d.Language
		}
	}
	if p.HardwareConcurrency == 0 {
		p.HardwareConcurrency = d.HardwareConcurrency
	}
	return p
}

type Navigator struct {
	entity.Entity
	profile NavigatorProfile
}

func (n *Navigator) UserAgent() string        { return n.profile.UserAgent }
func (n *Navigator) Platform() string         { return n.profile.Platform }
func (n *Navigator) Vendor() string           { return n.profile.Vendor }
func (n *Navigator) Language() string         { return n.profile.Language }
func (n *Navigator) Languages() []string      { return n.profile.Languages }
func (n *Navigator) HardwareConcurrency() int { return n.profile.HardwareConcurrency }

// Webdriver reports navigator.webdriver. gost-dom drives pages
// programmatically, but reports false to mirror an ordinary browsing session.
func (n *Navigator) Webdriver() bool { return false }
