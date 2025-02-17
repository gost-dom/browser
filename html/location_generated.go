// This file is generated. Do not edit.

package html

type Location interface {
	Href() string
	SetHref(string)
	Origin() string
	Protocol() string
	SetProtocol(string)
	Host() string
	SetHost(string)
	Hostname() string
	SetHostname(string)
	Port() string
	SetPort(string)
	Pathname() string
	SetPathname(string)
	Search() string
	SetSearch(string)
	Hash() string
	SetHash(string)
	AncestorOrigins() DOMStringList
	Assign(string) error
	Replace(string) error
	Reload() error
}
