package sobekhost

import (
	"reflect"
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/grafana/sobek"
)

const internal_symbol_name = "__go_dom_internal_value__"

func New() html.ScriptHost {
	return &gojaScriptHost{}
}

type gojaScriptHost struct{}

type propertyNameMapper struct{}

func (_ propertyNameMapper) FieldName(t reflect.Type, f reflect.StructField) string {
	return ""
}

func uncapitalize(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}

func (_ propertyNameMapper) MethodName(t reflect.Type, m reflect.Method) string {
	var doc dom.Document
	var document = reflect.TypeOf(&doc).Elem()
	if t.Implements(document) && m.Name == "Location" {
		return uncapitalize(m.Name)
	} else {
		return ""
	}
}

func (d *gojaScriptHost) NewContext(window html.Window) html.ScriptContext {
	vm := sobek.New()
	vm.SetFieldNameMapper(propertyNameMapper{})
	result := &GojaContext{
		vm:           vm,
		clock:        clock.New(),
		window:       window,
		wrappedGoObj: sobek.NewSymbol(internal_symbol_name),
		cachedNodes:  make(map[int32]sobek.Value),
		classes:      make(map[string]*gojaClass),
	}

	globalThis := vm.GlobalObject()
	globalThis.DefineDataPropertySymbol(
		result.wrappedGoObj,
		vm.ToValue(window),
		sobek.FLAG_FALSE,
		sobek.FLAG_FALSE,
		sobek.FLAG_FALSE,
	)
	globalThis.Set("window", globalThis)
	initializer.Configure(result)
	location := result.createLocationInstance()
	globalThis.DefineAccessorProperty(
		"location",
		vm.ToValue(func(c *sobek.FunctionCall) sobek.Value {
			return location
		}),
		nil,
		sobek.FLAG_FALSE,
		sobek.FLAG_TRUE,
	)
	globalThis.SetPrototype(result.classes["Window"].prototype)

	return result
}

func (d *gojaScriptHost) Close() {}
