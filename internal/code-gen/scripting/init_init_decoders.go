package scripting

import (
	"fmt"
	"os"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/code-gen/events"
	"github.com/gost-dom/code-gen/gen"
	"github.com/gost-dom/code-gen/internal"
	"github.com/gost-dom/code-gen/packagenames"
	"github.com/gost-dom/code-gen/stdgen"
	g "github.com/gost-dom/generators"
	"github.com/gost-dom/webref/idl"
)

func CreateEventInitDecoder(name string, dict idl.Dictionary) g.Generator {
	decoderName := fmt.Sprintf("decode%s", name)
	eventType := g.NewTypePackage(name, packagenames.UIEvents)

	return gen.NewFunction(
		gen.FunctionName(decoderName),
		gen.FunctionTypeParam(gen.AnyConstraint(g.Id("T"))),
		gen.FunctionParam(g.Id("scope"), jsScope),
		gen.FunctionParam(g.Id("options"), jsObject),
		gen.FunctionParam(g.Id("init"), eventType.Pointer()),
		gen.FunctionReturnType(g.Id("error")),
		gen.FunctionBody(createEventInitDecoderBody(dict)),
	)
}

func createEventInitDecoderBody(dict idl.Dictionary) g.Generator {
	var decoders []g.Generator
	var init = g.NewValue("init")
	var options = g.Id("options")
	if inheritance := dict.Inheritance; inheritance != "" {
		parentDecoder := g.NewValue(fmt.Sprintf("decode%s", inheritance))
		decoders = append(decoders, gen.NewlineBefore(parentDecoder.Call(
			g.Id("scope"),
			g.Id("options"),
			init.Field(inheritance).Reference(),
		)))
	}

	for _, entry := range dict.Entries {
		var fieldName = internal.IdlNameToGoName(entry.Key)
		var field = init.Field(fieldName)
		decoders = append(
			decoders,
			gen.NewlineBefore(jsDecodeInto.Call(
				g.Id("scope"),
				field.Reference(),
				options,
				g.Lit(entry.Key),
				decoderForType(entry.Value),
			),
			))
	}
	decoders = append(decoders, g.Line)

	return g.Return(stdgen.ErrorsJoin(decoders...))
}

func CreateEventInitDecoders(pkg string) error {
	spec, err := idl.Load(pkg)
	if err != nil {
		return fmt.Errorf("CreateEventInitDecoders: load pkg %s: %w", pkg, err)
	}
	statements := g.StatementList()
	for _, name := range events.GeneratedEventInitNames {
		dict, ok := spec.Dictionaries[name]
		if !ok {
			return fmt.Errorf("CreateEventInitDecoders: %s: dictionary not found", name)
		}
		statements.Append(g.Line, CreateEventInitDecoder(name, dict))
	}

	file := jen.NewFile(pkg)
	file.HeaderComment("This file is generated. Do not edit.")
	file.Add(statements.Generate())

	filename := "event_inits_generated.go"
	writer, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer writer.Close()
	return file.Render(writer)
}
