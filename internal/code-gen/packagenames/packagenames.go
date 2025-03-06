package packagenames

import "fmt"

const (
	NAME      = "gost-dom"
	BASE_PKG  = "github.com/gost-dom/browser"
	ISSUE_URL = "https://github.com/gost-dom/browser/issues"

	Dom      = BASE_PKG + "/dom"
	Events   = Dom + "/event"
	Html     = BASE_PKG + "/html"
	V8host   = BASE_PKG + "/scripting/v8host"
	Gojahost = BASE_PKG + "/scripting/gojahost"
	Log      = BASE_PKG + "/internal/log"
)

var names = map[string]string{
	"dom":            Dom,
	"html":           Html,
	"uievents":       Dom,
	"pointerevents4": Dom,
}

// PackageName returns the package name containing the implementation of a
// specific web API.
func PackageName(apiName string) string {
	if res, found := names[apiName]; found {
		return res
	}
	return fmt.Sprintf("%s/internal/%s", BASE_PKG, apiName)
}
