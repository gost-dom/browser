package mutation

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/gosterror"
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
)

type Record = dominterfaces.MutationRecord

type Callback interface {
	HandleMutation([]Record, *Observer)
}

type Flusher interface {
	Flush()
}

type Flushers interface {
	AddFlusher(Flusher)
	RemoveFlusher(Flusher)
}

type CallbackFunc func([]Record, *Observer)

func (f CallbackFunc) HandleMutation(r []Record, o *Observer) { f(r, o) }

type RecordCallbackFunc func([]Record)

func (f RecordCallbackFunc) HandleMutation(r []Record, _ *Observer) { f(r) }

// Observer implements behaviour of the [MutationObserver], allowing client code
// to react to changes to the DOM. Mutation events are not published
// immediately; but before control returns to the event loop, (i.e. if an event
// hander generates two mutation events, an observer will not receive any events
// before the function has finished executing). The records can be pulled from
// the [Observer.TakeRecords] function.
//
// Client code must first call [MutationObserver.Observe], specifying what to
// listen for.
//
// The Observer must be initialized with both a [Flushers], allowing it to
// register to receive notifications; and a Callback.
//
// [MutationObserver]: https://developer.mozilla.org/en-US/docs/Web/API/MutationObserver
type Observer struct {
	Flushers Flushers
	Callback Callback
	pending  []Record
	closer   dom.Closer
	options  Options
	target   dom.Node
}

// NewObserver creates a new observer registering with the [Flushers] f and the
// callback cb. NewObserver doesn't do anything except setting exported fields.
// The function panics if either f or cb are nil.
func NewObserver(f Flushers, cb Callback) *Observer {
	if f == nil {
		panic("mutation: NewObserver: f cannot be nil")
	}
	if cb == nil {
		panic("mutation: NewObserver: cb cannot be nil")
	}
	return &Observer{Flushers: f, Callback: cb}
}

// Start observing for changes for a specific dom node.
//
// Panics if the observer does not have a handler.
func (o *Observer) Observe(node dom.Node, options ...func(*Options)) error {
	o.Flushers.AddFlusher(o)
	o.assertCanObserve()

	o.options = Options{}
	for _, opt := range options {
		opt(&o.options)
	}
	valid := o.options.ChildList || o.options.Attributes || o.options.CharacterData
	if !valid {
		return gosterror.NewTypeError(
			"MutationObserver.observe: One of 'childList', 'attributes', 'characterData' must not be false",
		)
	}

	o.target = node
	o.closer = node.Observe(o)
	return nil
}

func (o Observer) assertCanObserve() {
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
}

func (o *Observer) Disconnect() {
	o.Flushers.RemoveFlusher(o)
	o.closer.Close()
	o.closer = nil
}

func (o *Observer) TakeRecords() (res []Record) {
	res = o.pending
	o.pending = nil
	return
}

func (o *Observer) Process(e dom.ChangeEvent) {
	if e.Target != o.target && !o.options.Subtree {
		return
	}
	switch e.Type {
	case dom.ChangeEventAttributes:
		if !o.options.Attributes {
			return
		}
	case dom.ChangeEventCData:
		if !o.options.CharacterData {
			return
		}
	case dom.ChangeEventChildList:
		if !o.options.ChildList {
			return
		}
	}
	r := Record{
		Type:         string(e.Type),
		Target:       e.Target,
		AddedNodes:   e.AddedNodes,
		RemovedNodes: e.RemovedNodes,
	}
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
