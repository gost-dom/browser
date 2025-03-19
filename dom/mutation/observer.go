package mutation

import (
	"github.com/gost-dom/browser/dom"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
)

type Record = dominterfaces.MutationRecord

type Callback interface {
	HandleMutation([]Record, *Observer)
}

type Options = dominterfaces.MutationObserverInit

type CallbackFunc func([]Record, *Observer)

func (f CallbackFunc) HandleMutation(r []Record, o *Observer) { f(r, o) }

type RecordCallbackFunc func([]Record)

func (f RecordCallbackFunc) HandleMutation(r []Record, _ *Observer) { f(r) }

type Observer struct {
	Callback Callback
	pending  []Record
	closer   dom.Closer
}

func NewObserver(cb Callback) *Observer {
	return &Observer{Callback: cb}
}

func (o *Observer) Observe(dom.Node) {
	panic("This function should be removed, empty options are invalid. Why it the value optional?")
}

// Start observing for changes for a specific dom node.
//
// Panics if the observer does not have a handler.
func (o *Observer) ObserveOptions(node dom.Node, options Options) {
	if o.Callback == nil {
		// Why panic and not ignore?
		//
		// We could easily ignore the call in flush, but if you don't set a
		// handler, you are using this type incorrectly in the first place.
		//
		// So we could panic in Flush instead. However, that makes the panic
		// more disconnected from the code that is flawed. The call stack would
		// originate from a DOM mutating call, such as Node.AppendChild; where
		// this is in a call stask where test code (assumably) is setting up an
		// observer.
		//
		// So panicing is the most helpful for a developer to find the issue in
		// client code.
		//
		// Client code can still _remove_ the handler after creating. But that
		// problem is not dealt with here, and will implicitly panic when
		// flushing.
		panic("Observer.ObserveOptions: A handler must be set before")
	}
	if o.closer != nil {
		panic("Observer.ObserveOptions: Observer is already observing a DOM node")
	}
	o.closer = node.Observe(o)
}

func (o *Observer) Disconnect() {
	o.closer.Close()
	o.closer = nil
}

func (o *Observer) TakeRecords() (res []Record) {
	res = o.pending
	o.pending = nil
	return
}

func (o *Observer) Process(e dom.ChangeEvent) {
	r := Record{Target: e.Target, AddedNodes: e.AddedNodes}
	o.pending = append(o.pending, r)
}

// Deprecated: Flush is a temporary solution while developing, and is not
// intended to be called by client code.
func (o *Observer) Flush() {
	records := o.TakeRecords()
	if len(records) > 0 {
		o.Callback.HandleMutation(records, o)
	}
}
