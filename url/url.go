package url

import (
	"fmt"
	netURL "net/url"
	"strings"
)

type URL struct {
	url *netURL.URL
}

func NewUrl(rawUrl string) (*URL, error) {
	if res, err := netURL.Parse(rawUrl); err == nil {
		return &URL{res}, nil
	} else {
		return nil, err
	}
}

func NewUrlBase(relativeUrl string, base string) (result *URL, err error) {
	rel, err := netURL.Parse(relativeUrl)
	if err == nil {
		if rel.Host != "" {
			return NewURLFromNetURL(rel), nil
		}
	}
	var u *netURL.URL
	if u, err = netURL.Parse(base); err != nil {
		return
	}
	u.RawQuery = ""
	u.Fragment = ""
	if strings.HasPrefix(relativeUrl, "/") {
		u.Path = rel.Path
		u.RawQuery = rel.RawQuery
		u.Fragment = rel.Fragment
	} else {
		// A DOM Url treats the relative path from the last slash in the base URL.
		// Go's URL doesn't. To get the right behaviour, use the parent dir if the
		// base path doesn't end in a slash
		if u.Path != "" && !strings.HasSuffix(u.Path, "/") {
			u = u.JoinPath("..")
		}
		u = u.JoinPath(relativeUrl)
	}
	result = NewURLFromNetURL(u)
	return
}

func ParseURL(rawUrl string) *URL {
	res, err := NewUrl(rawUrl)
	if err != nil {
		res = nil
	}
	return res
}

func ParseURLBase(relativeUrl string, base string) *URL {
	res, err := NewUrlBase(relativeUrl, base)
	if err != nil {
		res = nil
	}
	return res
}

func CanParseURL(rawUrl string) bool {
	_, err := NewUrl(rawUrl)
	return err == nil
}

func CreateObjectURL(object any) (*URL, error) {
	panic("Not implemented")
	// return nil, dom.newNotImplementedError("URL.CreateObjectURL not implemented yet")
}

func RevokeObjectURL(object any) (*URL, error) {
	panic("Not implemented")
	// return nil, dom.newNotImplementedError("URL.RevokeObjectURL not implemented yet")
}

func NewURLFromNetURL(u *netURL.URL) *URL {
	return &URL{u}
}

func (l URL) Hash() string {
	if l.url.Fragment == "" {
		return ""
	}
	return "#" + l.url.Fragment
}

func (l URL) Host() string { return l.url.Host }

func (l URL) Hostname() string {
	return l.url.Hostname()
}

func (l URL) String() string { return l.Href() }
func (l URL) Href() string   { return l.url.String() }

func (l URL) Origin() string { return l.url.Scheme + "://" + l.url.Host }

func (l URL) Pathname() string { return l.url.Path }

func (l URL) Protocol() string { return l.url.Scheme + ":" }

func (l URL) Port() string { return l.url.Port() }

func (l URL) Search() string {
	if l.url.RawQuery != "" {
		return "?" + l.url.RawQuery
	} else {
		return ""
	}
}

func (l URL) ToJSON() (string, error) { return l.Href(), nil }

func (l URL) Password() string { p, _ := l.url.User.Password(); return p }
func (l *URL) SetPassword(val string) {
	l.url.User = netURL.UserPassword(l.Username(), val)
}

func (l *URL) SetUsername(
	val string,
) {
	p, _ := l.url.User.Password()
	l.url.User = netURL.UserPassword(val, p)
}
func (l URL) Username() string { return l.url.User.Username() }
func (l URL) SearchParams() URLSearchParams {
	// TODO
	panic("URL.SearchParams Not implemented. Please file a feature request")
}

func (l *URL) SetHash(val string) { l.url.Fragment = val }
func (l *URL) SetHost(val string) { l.url.Host = val }
func (l *URL) SetHostname(val string) {
	port := l.Port()
	if port == "" {
		l.url.Host = val
	} else {
		l.url.Host = fmt.Sprintf("%s:%s", val, port)
	}
}
func (l *URL) SetHref(val string) {
	v, err := netURL.Parse(val)
	// TODO: Generate error
	if err != nil {
		panic(err)
	}
	l.url = v
}
func (l *URL) SetPathname(val string) { l.url.Path = val }
func (l *URL) SetPort(val string) {
	if val == "" {
		l.url.Host = l.url.Hostname()
	} else {
		l.url.Host = fmt.Sprintf("%s:%s", l.Hostname(), val)
	}
}
func (l *URL) SetProtocol(val string) { l.url.Scheme = val }
func (l *URL) SetSearch(val string) {
	if strings.HasPrefix(val, "?") {
		val = val[1:]
	}
	l.url.RawQuery = val
}
