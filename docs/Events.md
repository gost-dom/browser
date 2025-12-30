# Events

Events are modelled slightly differently in Go than in JavaScript

## JavaScript

In JavaScript, there's an inheritance hierarchy between events. `KeyboardEvent`
inherits from `UIEvent`, which inherits from `Event`.

A `KeyboardEvent` accepts a `KeyboardEventInit` dictionary, which "inherits"
`UIEventInit`. The types are not runtime types, merely specifications for valid
fields.

The properties on the "EventInit" object become properties on the event.

At runtime, you can check the type of the event, e.g. `event instanceof
KeyboardEvent`

## In Go

As Go doesn't handle the hierarchy very well, all events are the same type,
`events.Event`.

```Go
package events 
type Event struct {
	// More members, just for 
	Type       string
	Bubbles    bool
	Cancelable bool
	Data       any
}
```

The `Data` field contains `EventInit` types, which use embedding to reflect the
"inheritance" hierarchy.

```Go
type KeyboardEventInit struct {
    UIEventInit
    Key  string
    Code string
    // etc....
}
```

## Go <-> Js mapping

When a concrete event type is created in JavaScript, e.g., using `new
KeyboardEvent('type', {})`, the `Data` field is set to the corresponding
`EventInit` type.

When an event is exposed to JavaScript, a new JavaScript object is created with
the correct prototype by inspecting the run-time type of the `Data`. I.e., if
`Data` is a `KeyboardEventInit`, a JavaScript `KeyboardEvent` object is returned.

The prototype of each event type has accessor attributes; where the underlying
Go implementation expects an `*event.Event` value with a specific type for
`Data`. E.g., the JavaScript `KeyboardEvent` prototype has a `key` attribute,
and the getter implementation in Go expects a `KeyboardEventInit` value in the
`Data` field.
