package events

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/packagenames"
	gen "github.com/gost-dom/generators"
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

func (g EventInitGenerator) Generate() *jen.Statement {
	return gen.Assign(
		gen.NewValue("data"),
		g.EventType.InitType().CreateInstance(),
	).Generate()
}

// TODO: Delete
type EventPropertiesGenerator struct{ EventType }

func (g EventPropertiesGenerator) Generate() *jen.Statement {
	event := gen.NewValue("event")
	s := gen.StatementList()
	keys := make([]string, 0)
	for k := range g.Options {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := g.Options[events.EventOption(k)]
		field := internal.UpperCaseFirstLetter(string(k))
		value := gen.Lit(v)
		s.Append(gen.Reassign(event.Field(field), value))
	}
	return s.Generate()
}

// TODO: Delete
type EventConstructorGenerator struct{ EventType }

func (s EventConstructorGenerator) Generate() *jen.Statement {
	return gen.StatementList(
		gen.StructLiteral{
			Type: gen.NewTypePackage("Event", packagenames.Events),
			Elements: []gen.Generator{
				gen.StructLiteralKeyElement{Key: gen.Id("Type"), Value: gen.Lit(s.Type)},
				gen.StructLiteralKeyElement{Key: gen.Id("Data"), Value: gen.Id("data")},
			},
		}.Value().Reference(),
	).Generate()

}

type DispatchEventInitBodyGenerator struct {
	EventType
}

func (g DispatchEventInitBodyGenerator) Generate() *jen.Statement {
	receiver := gen.NewValue("e")
	event := gen.NewValue("event")
	return gen.StatementList(
		gen.Assign(event, EventConstructorGenerator(g)),
		EventPropertiesGenerator(g),
		gen.Return(
			receiver.Field("DispatchEvent").Call(event, gen.Line))).Generate()
}

// Renders the line to dispatch an event
//
//	data := PointerEventInit{}
//	// ...
//	return e.DispatchEvent(event)
type DispatchEventBodyGenerator struct {
	EventType
}

func (g DispatchEventBodyGenerator) Generate() *jen.Statement {
	receiver := gen.NewValue("e")
	event := gen.NewValue("event")
	return gen.StatementList(
		EventInitGenerator{g.EventType},
		gen.Assign(event, EventConstructorGenerator(g)),
		EventPropertiesGenerator(g),
		gen.Return(
			receiver.Field("DispatchEvent").Call(event, gen.Line))).Generate()
}

// Generates a single event dispatch methods for an element, e.g.
//
//	func Blur(e Element) { /* ... */ }
type EventDispatchMethodGenerator struct {
	Type       gen.Type
	TargetType EventTargetType
	Event      EventType
}

func eventDispatchTypeName(typeName string) string {
	return fmt.Sprintf("%sEvents", typeName)
}

func (g EventDispatchMethodGenerator) Generate() *jen.Statement {
	event := g.Event
	return gen.StatementList(
		DispatchFunctionDocumentation{EventType: g.Event.Type, Target: g.TargetType},
		gen.FunctionDefinition{
			Args:     gen.Arg(gen.Id("e"), g.TargetType),
			Name:     internal.UpperCaseFirstLetter(event.Type),
			RtnTypes: []gen.Generator{gen.Id("bool")},
			Body:     DispatchEventBodyGenerator{event},
		},
	).Generate()
}

// Generates a single event dispatch methods for an element with data, e.g.
//
//	func Keydown(e Element, data KeyboardEventInit) { /* ... */ }
type DispatchEventInitGenerator struct {
	Type       gen.Type
	TargetType EventTargetType
	Event      EventType
}

func (g DispatchEventInitGenerator) Generate() *jen.Statement {
	event := g.Event

	return gen.StatementList(
		DispatchFunctionDocumentation{EventType: g.Event.Type, Target: g.TargetType},
		gen.FunctionDefinition{
			Args:     gen.Arg(gen.Id("e"), g.TargetType).Arg(gen.Id("data"), g.Event.InitType()),
			Name:     internal.UpperCaseFirstLetter(event.Type) + "Init",
			RtnTypes: []gen.Generator{gen.Id("bool")},
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

func (g EventInterfaceGenerator) Generate() *jen.Statement {
	name := fmt.Sprintf("%sEvents", g.Element)
	ops := make([]gen.Generator, len(g.Events))
	for i, e := range g.Events {
		ops[i] = gen.Raw(
			jen.Commentf("Deprecated: %s is not a method defined on Element in the DOM", e.Type).
				Line().
				Id(internal.UpperCaseFirstLetter(e.Type)).
				Params().
				Add(jen.Id("bool")),
		)
	}
	return jen.Commentf("Deprecated: %s expose methods that are not part of the %s specification ", name, g.Element).
		Line().
		Type().
		Add(jen.Id(name)).
		Interface(gen.ToJenCodes(ops)...)
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

func (g ElementEventGenerator) Type() gen.Type {
	return gen.NewType(eventDispatchTypeName(g.TargetType.String()))
}

// EventType represents a specific event type specification, and has helper
// functions, e.g. to generate the corresponding Go data types.
type EventType events.Event

func (t EventType) InitType() gen.Type {
	return gen.NewType(fmt.Sprintf("%sInit", t.Interface))
}

func (g ElementEventGenerator) Generate() *jen.Statement {
	events := g.Events()
	type_ := g.Type()

	res := gen.StatementList(
		// ElementEventStructGenerator{type_},
		// gen.Line,
		// EventInterfaceGenerator{Element: g.Element, Events: events},
		// gen.Line,
		EventDispatchMethodsGenerator{
			Type:       type_,
			TargetType: g.TargetType,
			Events:     events,
		},
	)
	return res.Generate()
}

func (g ElementEventGenerator) Events() []EventType {
	var events []EventType
	for _, e := range g.Api.EventsForType(g.TargetType.String()) {
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
	return gen.NewTypePackage(t.TypeName, t.Package).Generate()
}

// Generates
type ElementEventStructGenerator struct {
	Type gen.Type
}

func (g ElementEventStructGenerator) Generate() *jen.Statement {
	s := gen.Struct{Name: g.Type}
	s.Field(gen.Id("Target"), gen.NewTypePackage("EventTarget", packagenames.Events))
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
	Type       gen.Type
	TargetType EventTargetType
	Events     []EventType
}

func (g EventDispatchMethodsGenerator) Generate() *jen.Statement {
	res := gen.StatementList()
	for _, e := range g.Events {
		res.Append(gen.Line)
		res.Append(EventDispatchMethodGenerator{
			Type:       g.Type,
			TargetType: g.TargetType,
			Event:      e,
		})
		res.Append(gen.Line)
		res.Append(DispatchEventInitGenerator{
			Type:       g.Type,
			TargetType: g.TargetType,
			Event:      e,
		})
	}
	return res.Generate()
}

func CreateEventSourceGenerator(apiName string, element string) (gen.Generator, error) {
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

type ElementEventDispatchFunctionsGenerator struct {
	Api     events.Events
	Element string
}

func (g *ElementEventDispatchFunctionsGenerator) Generate() *jen.Statement {
	return nil
}
