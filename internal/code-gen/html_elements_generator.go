package main

import (
	"fmt"
	"io"

	"github.com/gost-dom/webref/elements"
)

func WriteHeader(b *builder, pkg string) {
	b.Printf("// This file is generated. Do not edit.\n\n")
	b.Printf("package %s\n\n", pkg)
}

type builder struct {
	io.Writer
	indentLvl int
}

func newBuilder(w io.Writer) *builder {
	return &builder{w, 0}
}

func (b builder) Printf(format string, args ...any) {
	for range b.indentLvl {
		fmt.Fprint(b.Writer, "\t")
	}
	fmt.Fprintf(b.Writer, format, args...)
}

func (b *builder) indent() {
	b.indentLvl++
}

func (b *builder) indentF(format string, args ...interface{}) {
	b.indent()
	b.Printf(format, args...)
}

func (b *builder) unIndent() {
	b.indentLvl--
	if b.indentLvl < 0 {
		panic("More unindentation than indentation")
	}
}

func (b *builder) unIndentF(format string, args ...interface{}) {
	b.unIndent()
	b.Printf(format, args...)
}

func generateHtmlElements(writer io.Writer) error {
	e, err := elements.Load("html")
	if err == nil {
		file := newBuilder(writer)
		WriteHeader(file, "codec")
		fmt.Fprint(file, "var HtmlElements = map[string]string {\n")
		file.indent()
		defer file.unIndentF("}\n")
		for _, element := range e.Elements {
			file.Printf("\"%s\": \"%s\",\n", element.Name, element.Interface)
		}
	}
	return err
}
