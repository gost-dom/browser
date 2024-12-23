package main

import (
	_ "embed"

	"github.com/dave/jennifer/jen"
)

//go:embed webref/curated/idlparsed/dom.json
var domData []byte

func generateDOMTypes(b *builder) error {
	file := jen.NewFilePath(sc)
	file.HeaderComment("This file is generated. Do not edit.")
	file.ImportName(br, "browser")
	file.ImportAlias(v8, "v8")
	data, err := createData(domData, CreateDataData{
		TypeName: "DOMTokenList",
		Receiver: "u",
		Customization: []string{
			"item",
			"contains",
			"remove",
			"toggle",
			"replace",
			"supports",
			"value",
		},
	})
	if err != nil {
		return err
	}
	writeFactory(file, data)
	return file.Render(b)
}
