package events

import (
	"fmt"
	"os"
	"slices"
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
	init := gen.NewValue("init")
	typeName := fmt.Sprintf("%sInit", g.Interface)
	s := gen.StatementList(gen.Assign(init, gen.Raw(jen.Id(typeName).Values())))
	keys := make([]string, 0)
	for k := range g.Options {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := g.Options[events.EventOption(k)]
		field := internal.UpperCaseFirstLetter(string(k))
		value := gen.Lit(v)
		s.Append(gen.Reassign(init.Field(field), value))
	}
	return s.Generate()
}

type EventConstructorGenerator events.Event

func (s EventConstructorGenerator) Generate() *jen.Statement {
	e := events.Event(s)
	eventConstructor := fmt.Sprintf("New%s", e.Interface)
	// arguments := []gen.Generator{
	// 	gen.Lit(e.Type),
	// }
	// if b, ok := e.Options["bubbles"]; ok {
	// 	arguments = append(
	// 		arguments,
	// 		gen.NewValuePackage("EventBubbles", packagenames.Events).Call(gen.Lit(b)),
	// 	)
	// }
	// if b, ok := e.Options["cancelable"]; ok {
	// 	arguments = append(
	// 		arguments,
	// 		gen.NewValuePackage("EventCancelable", packagenames.Events).Call(gen.Lit(b)),
	// 	)
	// }
	// if b, ok := e.Options["composable"]; ok {
	// 	// This is theoretical. There are no composable event
	// 	// definitions in the source data.
	// 	arguments = append(arguments, gen.NewValue("EventComposable").Call(gen.Lit(b)))
	// }
	return gen.StatementList(
		gen.NewValue(eventConstructor).Call(gen.Lit(e.Type), gen.NewValue("init")),
	).Generate()

}

type DispatchEventGenerator events.Event

func (g DispatchEventGenerator) Generate() *jen.Statement {
	e := events.Event(g)
	receiver := gen.NewValue("e")
	return gen.StatementList(
		EventInitGenerator(g),
		gen.Return(
			receiver.Field("target").Field("DispatchEvent").Call(
				line(EventConstructorGenerator(e)), gen.Line))).Generate()

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

func CreateEventSourceGenerator(apiName string, element string) (gen.Generator, error) {
	api, err := events.Load(apiName)
	n := gen.NewType(eventDispatchTypeName(element))
	s := gen.Struct{Name: n}
	s.Field(gen.Id("target"), gen.NewTypePackage("EventTarget", packagenames.Events))
	events := slices.DeleteFunc(
		api.EventsForType(element),
		func(e events.Event) bool { return e.Interface != "PointerEvent" },
	)
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
