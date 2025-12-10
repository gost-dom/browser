package matchers

import (
	"errors"

	"github.com/gost-dom/browser/dom"

	"github.com/onsi/gomega/gcustom"
	. "github.com/onsi/gomega/types"
)

func BeADOMError() GomegaMatcher {
	return gcustom.MakeMatcher(func(e error) (bool, error) {
		return errors.Is(e, dom.ErrDom), nil
	})
}
