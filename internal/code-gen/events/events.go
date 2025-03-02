package events

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/internal"
	gen "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/events"
)

type EventGeneratorSpecs struct {
	Api        string
	SourceType string
	EventName  string
}

type EventConstructorGenerator events.Event

func (s EventConstructorGenerator) Generate() *jen.Statement {
	e := events.Event(s)
	eventConstructor := fmt.Sprintf("New%s", e.Interface)
	arguments := []gen.Generator{
		gen.Lit(e.Type),
	}
	if b, ok := e.Options["bubbles"]; ok {
		arguments = append(arguments, gen.NewValue("EventBubbles").Call(gen.Lit(b)))
	}
	if b, ok := e.Options["cancelable"]; ok {
		arguments = append(arguments, gen.NewValue("EventCancelable").Call(gen.Lit(b)))
	}
	if b, ok := e.Options["composable"]; ok {
		// This is theoretical. There are no composable event
		// definitions in the source data.
		arguments = append(arguments, gen.NewValue("EventComposable").Call(gen.Lit(b)))
	}
	return gen.NewValue(eventConstructor).Call(arguments...).Generate()

}

type DispatchEventGenerator events.Event

func (g DispatchEventGenerator) Generate() *jen.Statement {
	e := events.Event(g)
	receiver := gen.NewValue("e")
	return gen.Return(
		receiver.Field("target").Field("DispatchEvent").Call(
			line(EventConstructorGenerator(e)), gen.Line)).Generate()

}

type GetGeneratorsRes struct {
	optionsGenerator gen.Generator
}

func line(g gen.Generator) gen.Generator { return gen.Raw(jen.Line().Add(g.Generate())) }

type EventDispatchMethodGenerator struct {
	SourceTypeName string
	Event          events.Event
}

func (g EventDispatchMethodGenerator) Generate() *jen.Statement {
	typeName := fmt.Sprintf("%sEvents", internal.LowerCaseFirstLetter(g.SourceTypeName))
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

func CreateEventSourceGenerator(api string, element string) (gen.Generator, error) {
	return nil, nil
}
