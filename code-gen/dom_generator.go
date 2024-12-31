package main

import (
	_ "embed"

	"github.com/dave/jennifer/jen"
	g "github.com/stroiman/go-dom/code-gen/generators"
)

//go:embed webref/curated/idlparsed/dom.json
var domData []byte

//go:embed webref/curated/idlparsed/html.json
var htmlData []byte

func generateDOMTypes(b *builder) error {
	file := jen.NewFilePath(sc)
	file.HeaderComment("This file is generated. Do not edit.")
	file.ImportName(br, "browser")
	file.ImportAlias(v8, "v8")

	domTokenList := ESClassWrapper{
		TypeName:      "DOMTokenList",
		Receiver:      "u",
		RunCustomCode: true,
	}
	domTokenList.Method("item").SetNoError()
	domTokenList.Method("contains").SetNoError()
	domTokenList.Method("remove").SetNoError()
	domTokenList.Method("toggle").SetCustomImplementation()
	domTokenList.Method("replace").SetNoError()
	domTokenList.Method("supports").SetNotImplemented()

	htmlTemplateElement := ESClassWrapper{
		TypeName: "HTMLTemplateElement",
		Receiver: "e",
	}
	htmlTemplateElement.Method("shadowRootMode").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootDelegatesFocus").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootClonable").SetNotImplemented()
	htmlTemplateElement.Method("shadowRootSerializable").SetNotImplemented()

	domTokenListData := createData(domData, domTokenList)
	htmlTemplateData := createData(htmlData, htmlTemplateElement)

	WriteGenerator(file, StatementList(
		domTokenListData,
		g.Line,
		htmlTemplateData,
	))
	return file.Render(b)
}
