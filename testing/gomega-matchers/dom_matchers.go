package matchers

import (
	"github.com/gost-dom/browser/dom"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gcustom"
	. "github.com/onsi/gomega/types"
)

func HaveAttribute(name string, expected interface{}) GomegaMatcher {
	data := struct {
		Attribute string
		Found     bool
		Actual    string
		Matcher   GomegaMatcher
		IsMatcher bool
	}{
		Attribute: name,
	}

	if expected == nil {
		return gcustom.MakeMatcher(func(e dom.Element) (bool, error) {
			_, found := e.GetAttribute(name)
			return found, nil
		}).WithTemplate(`Expected:\n{{.FormattedActual}}\n{{.To}} have attribute '{{.Data.Attribute}}'`, &data)
	}

	if data.Matcher, data.IsMatcher = expected.(GomegaMatcher); !data.IsMatcher {
		return HaveAttribute(name, gomega.Equal(expected))
	}
	return gcustom.MakeMatcher(func(e dom.Element) (bool, error) {
		if data.Actual, data.Found = e.GetAttribute(name); !data.Found {
			return false, nil
		}
		return data.Matcher.Match(data.Actual)
	}).WithTemplate(`Expected:\n{{.FormattedActual}}\n{{.To}} have attribute '{{.Data.Attribute}}'
{{ if .Data.Found}}{{.Data.Matcher.FailureMessage .Data.Actual -}}
{{else}}  Attribute did not exist{{end}}`, &data)
}
