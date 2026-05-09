// This file is generated. Do not edit.

package dominterfaces

import (
	"fmt"
	dom "github.com/gost-dom/browser/dom"
)

type Range interface {
	AbstractRange
	fmt.Stringer
	SetStart(dom.Node, int) error
	SetEnd(dom.Node, int) error
	Detach()
}
