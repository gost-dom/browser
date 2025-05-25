package packagenames

import "fmt"

// Names that relate to the project name, which would only change if the project
// is moved/renamed.
const (
	NAME      = "gost-dom"
	BASE_PKG  = "github.com/gost-dom/browser"
	ISSUE_URL = "https://github.com/gost-dom/browser/issues"
	// alternate url: "https://github.com/gost-dom/browser/issues?q=sort%3Aupdated-desc+is%3Aissue+is%3Aopen"

	Dom            = BASE_PKG + "/dom"
	Events         = Dom + "/event"
	Html           = BASE_PKG + "/html"
	V8host         = BASE_PKG + "/scripting/v8host"
	Gojahost       = BASE_PKG + "/scripting/gojahost"
	Log            = BASE_PKG + "/internal/log"
	UIEvents       = BASE_PKG + "/internal/uievents"
	DomInterfaces  = BASE_PKG + "/internal/interfaces/dom-interfaces"
	URLInterfaces  = BASE_PKG + "/internal/interfaces/url-interfaces"
	HTMLInterfaces = BASE_PKG + "/internal/interfaces/html-interfaces"
	HTMLInternal   = BASE_PKG + "/internal/html"
	JSAbstraction  = BASE_PKG + "/scripting/v8host/internal/abstraction"
	JS             = BASE_PKG + "/scripting/internal/js"
	URL            = BASE_PKG + "/url"
	V8go           = "github.com/gost-dom/v8go"
	Goja           = "github.com/dop251/goja"
)

var names = map[string]string{
	"dom":            Dom,
	"html":           Html,
	"pointerevents4": UIEvents,
}

// PackageName returns the package name containing the implementation of a
// specific web API.
func PackageName(apiName string) string {
	if res, found := names[apiName]; found {
		return res
	}
	return fmt.Sprintf("%s/internal/%s", BASE_PKG, apiName)
}
