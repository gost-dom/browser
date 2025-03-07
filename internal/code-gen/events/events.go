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

type EventInitGenerator events.Event

func (g EventInitGenerator) Generate() *jen.Statement {
	return gen.Assign(
		gen.NewValue("data"),
		gen.NewType(fmt.Sprintf("%sInit", g.Interface)).CreateInstance(),
	).Generate()
}

type EventPropertiesGenerator events.Event

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

type EventConstructorGenerator events.Event

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

type DispatchEventGenerator events.Event

func (g DispatchEventGenerator) Generate() *jen.Statement {
	e := events.Event(g)
	receiver := gen.NewValue("e")
	event := gen.NewValue("event")
	return gen.StatementList(
		EventInitGenerator(g),
		gen.Assign(event, EventConstructorGenerator(e)),
		EventPropertiesGenerator(g),
		gen.Return(
			receiver.Field("target").Field("DispatchEvent").Call(event, gen.Line))).Generate()
}

type GetGeneratorsRes struct {
	optionsGenerator gen.Generator
}

func line(g gen.Generator) gen.Generator { return gen.Raw(jen.Line().Add(g.Generate())) }

// Generates a single event dispatch methods for an element, e.g.
//
//	func (e *elementEvents) Blur() { /* ... */ }
type EventDispatchMethodGenerator struct {
	Type           gen.Type
	SourceTypeName string
	Event          EventType
}

func eventDispatchTypeName(typeName string) string {
	return fmt.Sprintf("%sEvents", internal.LowerCaseFirstLetter(typeName))
}

func (g EventDispatchMethodGenerator) Generate() *jen.Statement {
	event := g.Event
	return gen.FunctionDefinition{
		Receiver: gen.FunctionArgument{
			Name: gen.Id("e"),
			Type: g.Type.Pointer(),
		},
		Name:     internal.UpperCaseFirstLetter(event.Type),
		RtnTypes: []gen.Generator{gen.Id("bool")},
		Body:     DispatchEventGenerator(event),
	}.Generate()
}

type EventInterfaceGenerator struct {
	Element string
	Events  []EventType
}

func (g EventInterfaceGenerator) Generate() *jen.Statement {
	name := fmt.Sprintf("%sEvents", g.Element)
	ops := make([]gen.Generator, len(g.Events))
	for i, e := range g.Events {
		ops[i] = gen.Raw(jen.Id(internal.UpperCaseFirstLetter(e.Type)).Params().Add(jen.Id("bool")))
	}
	return jen.Type().Add(jen.Id(name)).Interface(gen.ToJenCodes(ops)...)
}

func IncludeEvent(e events.Event) bool {
	switch e.Interface {
	case "PointerEvent":
		return true
	case "FocusEvent":
		return !strings.HasPrefix(e.Type, "DOMFocus")
	default:
		return false
	}
}

// Generates the file containins event constructors for a specific element
//
//	type elementEvents struct { /* */ }
//
//	type ElementEvents interface { /* */ }
//
//	func (e *elementEvents) Click() { /* dispatch a "click" event */ }
type ElementEventGenerator struct {
	Api     events.Events
	Element string
}

func (g ElementEventGenerator) Type() gen.Type {
	return gen.NewType(eventDispatchTypeName(g.Element))
}

// EventType represents a specific event type specification, and has helper
// functions, e.g. to generate the corresponding Go data types.
type EventType events.Event

func (g ElementEventGenerator) Generate() *jen.Statement {
	events := g.Events()
	type_ := g.Type()

	res := gen.StatementList(
		ElementEventStructGenerator{type_},
		gen.Line,
		EventInterfaceGenerator{Element: g.Element, Events: events},
		gen.Line,
		EventDispatchMethodsGenerator{
			Type:           type_,
			SourceTypeName: g.Element,
			Events:         events,
		},
	)
	return res.Generate()
}

func (g ElementEventGenerator) Events() []EventType {
	var events []EventType
	for _, e := range g.Api.EventsForType(g.Element) {
		if IncludeEvent(e) {
			events = append(events, EventType(e))
		}
	}
	return events
}

// Generates
type ElementEventStructGenerator struct {
	Type gen.Type
}

func (g ElementEventStructGenerator) Generate() *jen.Statement {
	s := gen.Struct{Name: g.Type}
	s.Field(gen.Id("target"), gen.NewTypePackage("EventTarget", packagenames.Events))
	return s.Generate()
}

// Generates all event dispatch methods for an element, e.g. (but with blank
// lines between each.
//
//	func (e *elementEvents) Blur() { /* ... */ }
//	func (e *elementEvents) Click() { /* ... */ }
//	func (e *elementEvents) Focus() { /* ... */ }
type EventDispatchMethodsGenerator struct {
	Type           gen.Type
	SourceTypeName string
	Events         []EventType
}

func (g EventDispatchMethodsGenerator) Generate() *jen.Statement {
	res := gen.StatementList()
	for _, e := range g.Events {
		res.Append(gen.Line)
		res.Append(EventDispatchMethodGenerator{
			Type:           g.Type,
			SourceTypeName: g.SourceTypeName,
			Event:          e,
		})
	}
	return res.Generate()
}

func CreateEventSourceGenerator(apiName string, element string) (gen.Generator, error) {
	api, err := events.Load(apiName)
	if err != nil {
		return nil, err
	}
	return ElementEventGenerator{api, element}, nil
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
	"dom": {{
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
