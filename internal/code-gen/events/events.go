package events

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/packagenames"

	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/events"
)

type EventGeneratorSpecs struct {
	Api        string
	SourceType string
	EventName  string
}

// Renders the line to initialise the event init data:
//
//	data := PointerEventInit{}
type EventInitGenerator struct {
	EventType
}

func (ge EventInitGenerator) Generate() *jen.Statement {
	return g.Assign(
		g.NewValue("data"),
		ge.EventType.InitType().CreateInstance(),
	).Generate()
}

// TODO: Delete
type EventPropertiesGenerator struct{ EventType }

func (ge EventPropertiesGenerator) Generate() *jen.Statement {
	event := g.NewValue("event")
	s := g.StatementList()
	keys := make([]string, 0)
	for k := range ge.Options {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := ge.Options[events.EventOption(k)]
		field := internal.UpperCaseFirstLetter(string(k))
		value := g.Lit(v)
		s.Append(g.Reassign(event.Field(field), value))
	}
	return s.Generate()
}

// TODO: Delete
type EventConstructorGenerator struct{ EventType }

func (s EventConstructorGenerator) Generate() *jen.Statement {
	return g.StatementList(
		g.StructLiteral{
			Type: g.NewTypePackage("Event", packagenames.Events),
			Elements: []g.Generator{
				g.StructLiteralKeyElement{Key: g.Id("Type"), Value: g.Lit(s.Type)},
				g.StructLiteralKeyElement{Key: g.Id("Data"), Value: g.Id("data")},
			},
		}.Value().Reference(),
	).Generate()

}

type DispatchEventInitBodyGenerator struct {
	EventType
}

func (ge DispatchEventInitBodyGenerator) Generate() *jen.Statement {
	receiver := g.NewValue("e")
	event := g.NewValue("event")
	return g.StatementList(
		g.Assign(event, EventConstructorGenerator(ge)),
		EventPropertiesGenerator(ge),
		g.Return(
			receiver.Field("DispatchEvent").Call(event, g.Line))).Generate()
}

// Renders the line to dispatch an event
//
//	data := PointerEventInit{}
//	// ...
//	return e.DispatchEvent(event)
type DispatchEventBodyGenerator struct {
	EventType
}

func (ge DispatchEventBodyGenerator) Generate() *jen.Statement {
	receiver := g.NewValue("e")
	event := g.NewValue("event")
	return g.StatementList(
		EventInitGenerator{ge.EventType},
		g.Assign(event, EventConstructorGenerator(ge)),
		EventPropertiesGenerator(ge),
		g.Return(
			receiver.Field("DispatchEvent").Call(event, g.Line))).Generate()
}

// Generates a single event dispatch methods for an element, e.g.
//
//	func Blur(e Element) { /* ... */ }
type EventDispatchMethodGenerator struct {
	Type       g.Type
	TargetType EventTargetType
	Event      EventType
}

func eventDispatchTypeName(typeName string) string {
	return fmt.Sprintf("%sEvents", typeName)
}

func (ge EventDispatchMethodGenerator) Generate() *jen.Statement {
	event := ge.Event
	return g.StatementList(
		DispatchFunctionDocumentation{EventType: ge.Event.Type, Target: ge.TargetType},
		g.FunctionDefinition{
			Args:     g.Arg(g.Id("e"), ge.TargetType),
			Name:     internal.UpperCaseFirstLetter(event.Type),
			RtnTypes: []g.Generator{g.Id("bool")},
			Body:     DispatchEventBodyGenerator{event},
		},
	).Generate()
}

// Generates a single event dispatch methods for an element with data, e.g.
//
//	func Keydown(e Element, data KeyboardEventInit) { /* ... */ }
type DispatchEventInitGenerator struct {
	Type       g.Type
	TargetType EventTargetType
	Event      EventType
}

func (ge DispatchEventInitGenerator) Generate() *jen.Statement {
	event := ge.Event

	return g.StatementList(
		DispatchFunctionDocumentation{EventType: ge.Event.Type, Target: ge.TargetType},
		g.FunctionDefinition{
			Args:     g.Arg(g.Id("e"), ge.TargetType).Arg(g.Id("data"), ge.Event.InitType()),
			Name:     internal.UpperCaseFirstLetter(event.Type) + "Init",
			RtnTypes: []g.Generator{g.Id("bool")},
			Body:     DispatchEventInitBodyGenerator{event},
		},
	).Generate()
}

// Generates the code documentation for an event dispatch function
type DispatchFunctionDocumentation struct {
	// The event type, like "click", "blur", etc.
	EventType string
	Target    EventTargetType
}

func (d DispatchFunctionDocumentation) Generate() *jen.Statement {
	return jen.Commentf(
		`// Dispatches a [%s event]. Returns the return value from [EventTarget.DispatchEvent].
//
// The behaviour dictating the return value depends on the type of event. For
// more information see the [MDN docs]
//
// [%s event]: https://developer.mozilla.org/en-US/docs/Web/API/%s/%s_event
// [MDN docs]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent#return_value`,
		d.EventType,
		d.EventType,
		d.Target,
		d.EventType,
	)
}

type EventInterfaceGenerator struct {
	Element string
	Events  []EventType
}

func (ge EventInterfaceGenerator) Generate() *jen.Statement {
	name := fmt.Sprintf("%sEvents", ge.Element)
	ops := make([]g.Generator, len(ge.Events))
	for i, e := range ge.Events {
		ops[i] = g.Raw(
			jen.Commentf("Deprecated: %s is not a method defined on Element in the DOM", e.Type).
				Line().
				Id(internal.UpperCaseFirstLetter(e.Type)).
				Params().
				Add(jen.Id("bool")),
		)
	}
	return jen.Commentf("Deprecated: %s expose methods that are not part of the %s specification ", name, ge.Element).
		Line().
		Type().
		Add(jen.Id(name)).
		Interface(g.ToJenCodes(ops)...)
}

func IncludeEvent(e events.Event) bool {
	switch e.Interface {
	case "PointerEvent", "KeyboardEvent", "InputEvent":
		return true
	case "FocusEvent":
		return !strings.HasPrefix(e.Type, "DOMFocus")
	default:
		return false
	}
}

// Generates the file containins event dispatch methods for a specific element
//
//	func Click(e Element) { /* dispatch a "click" event */ }
//
//	func Submit(e HTMLFormElement) { /* dispatch a "submit" event for a form */ }
type ElementEventGenerator struct {
	Api        events.Events
	TargetType EventTargetType
}

func (ge ElementEventGenerator) Type() g.Type {
	return g.NewType(eventDispatchTypeName(ge.TargetType.String()))
}

// EventType represents a specific event type specification, and has helper
// functions, e.g. to generate the corresponding Go data types.
type EventType events.Event

func (t EventType) InitType() g.Type {
	return g.NewType(fmt.Sprintf("%sInit", t.Interface))
}

func (ge ElementEventGenerator) Generate() *jen.Statement {
	events := ge.Events()
	type_ := ge.Type()

	res := g.StatementList(
		// ElementEventStructGenerator{type_},
		// g.Line,
		// EventInterfaceGenerator{Element: g.Element, Events: events},
		// g.Line,
		EventDispatchMethodsGenerator{
			Type:       type_,
			TargetType: ge.TargetType,
			Events:     events,
		},
	)
	return res.Generate()
}

func (ge ElementEventGenerator) Events() []EventType {
	var events []EventType
	for _, e := range ge.Api.EventsForType(ge.TargetType.String()) {
		if IncludeEvent(e) {
			events = append(events, EventType(e))
		}
	}
	return events
}

// The specific EventTarget type, e.g., Element or HTMLFormElement
type EventTargetType struct {
	TypeName string
	Package  string
}

func (t EventTargetType) String() string { return t.TypeName }

func (t EventTargetType) Generate() *jen.Statement {
	return g.NewTypePackage(t.TypeName, t.Package).Generate()
}

// Generates
type ElementEventStructGenerator struct {
	Type g.Type
}

func (ge ElementEventStructGenerator) Generate() *jen.Statement {
	s := g.Struct{Name: ge.Type}
	s.Field(g.Id("Target"), g.NewTypePackage("EventTarget", packagenames.Events))
	return s.Generate()
}

// Generates all event dispatch methods for an element, e.g. (but with blank
// lines between each.
//
//	func Blur(e dom.Element) { /* ... */ }
//	func Click(e dom.Element) { /* ... */ }
//	func Focus(e dom.Element) { /* ... */ }
//
// The argument reflects the type that can be the source of the event. E.g., for
// the "submit" event, the argument type is HTMLFormElement.
type EventDispatchMethodsGenerator struct {
	Type       g.Type
	TargetType EventTargetType
	Events     []EventType
}

func (ge EventDispatchMethodsGenerator) Generate() *jen.Statement {
	res := g.StatementList()
	for _, e := range ge.Events {
		res.Append(g.Line)
		res.Append(EventDispatchMethodGenerator{
			Type:       ge.Type,
			TargetType: ge.TargetType,
			Event:      e,
		})
		res.Append(g.Line)
		res.Append(DispatchEventInitGenerator{
			Type:       ge.Type,
			TargetType: ge.TargetType,
			Event:      e,
		})
	}
	return res.Generate()
}

func CreateEventSourceGenerator(apiName string, element string) (g.Generator, error) {
	api, err := events.Load(apiName)
	if err != nil {
		return nil, err
	}
	return ElementEventGenerator{api, EventTargetType{element, packagenames.Dom}}, nil
}

func generateFile(packageName string, apiName string, element string) (*jen.File, error) {
	file := jen.NewFile(packageName)
	file.HeaderComment("This file is generated. Do not edit.")
	g, err := CreateEventSourceGenerator(apiName, element)
	if err == nil {
		file.Add(g.Generate())
	}
	return file, err
}

type eventSources struct {
	api   string
	names []string
}

var types = map[string][]eventSources{
	"uievents": {{
		api:   "uievents",
		names: []string{"Element"},
	}},
}

func CreateEventGenerators(packageName string) error {
	for _, source := range types[packageName] {
		for _, e := range source.names {
			var f *jen.File
			filename := fmt.Sprintf("%s_events_generated.go", strings.ToLower(e))
			writer, err := os.Create(filename)
			if err != nil {
				return err
			}
			defer writer.Close()
			f, err = generateFile(packageName, source.api, e)
			if err != nil {
				return err
			}
			f.Render(writer)
		}
	}
	return nil
}
