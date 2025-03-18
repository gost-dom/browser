package codegentest

import (
	"bytes"
	"io"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/gost-dom/generators"
)

type Renderer interface {
	Render(io.Writer) error
}

func RenderString(t testing.TB, renderer Renderer) string {
	t.Helper()
	var b bytes.Buffer
	err := renderer.Render(&b)
	if err != nil {
		t.Errorf("RenderString: Render returned: %s", err.Error())
	}
	return b.String()
}

func RenderInPackage(t testing.TB, packagePath string, gen generators.Generator) string {
	t.Helper()
	file := jen.NewFilePath(packagePath)
	file.Add(gen.Generate())
	return RenderString(t, file)
}
