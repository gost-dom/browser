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

type EventDispatchMethodGenerator struct {
	SourceTypeName string
	Event          events.Event
}

func eventDispatchTypeName(typeName string) string {
	return fmt.Sprintf("%sEvents", internal.LowerCaseFirstLetter(typeName))
}

func (g EventDispatchMethodGenerator) Generate() *jen.Statement {
	typeName := eventDispatchTypeName(g.SourceTypeName)
	event := g.Event
	return gen.FunctionDefinition{
		Receiver: gen.FunctionArgument{
			Name: gen.Id("e"),
			Type: gen.NewType(typeName).Pointer(),
		},
		Name:     internal.UpperCaseFirstLetter(event.Type),
		RtnTypes: []gen.Generator{gen.Id("bool")},
		Body:     DispatchEventGenerator(event),
	}.Generate()
}

func CreateMethodGenerator(specs EventGeneratorSpecs) (res gen.Generator, err error) {
	api, err := events.Load(specs.Api)
	events := api.EventsForType(specs.SourceType)
	for _, e := range events {
		if e.Type == specs.EventName {
			res = EventDispatchMethodGenerator{
				SourceTypeName: specs.SourceType,
				Event:          e,
			}
			break
		}
	}
	return
}

type EventInterfaceGenerator struct {
	Element string
	Events  []events.Event
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

func CreateEventSourceGenerator(apiName string, element string) (gen.Generator, error) {
	api, err := events.Load(apiName)
	n := gen.NewType(eventDispatchTypeName(element))
	s := gen.Struct{Name: n}
	s.Field(gen.Id("target"), gen.NewTypePackage("EventTarget", packagenames.Events))
	var events []events.Event
	for _, e := range api.EventsForType(element) {
		if IncludeEvent(e) {
			events = append(events, e)
		}
	}
	res := gen.StatementList(s, gen.Line, EventInterfaceGenerator{Element: element, Events: events})
	for _, e := range events {
		res.Append(gen.Line)
		res.Append(EventDispatchMethodGenerator{
			SourceTypeName: element,
			Event:          e,
		})
	}
	return res, err
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
