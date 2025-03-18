package mutation

import (
	"github.com/gost-dom/browser/dom"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
)

type Record = dominterfaces.MutationRecord

type Observer struct{}

func (o *Observer) Observe(dom.Node) {}

func (o *Observer) ObserveOptions(dom.Node, dominterfaces.MutationObserverInit) {}

func (o *Observer) Disconnect() {}

func (o *Observer) TakeRecords() []Record { return nil }
