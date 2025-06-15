package gojahost

import (
	"reflect"
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"

	"github.com/dop251/goja"
)

const internal_symbol_name = "__go_dom_internal_value__"

func New() html.ScriptHost {
	return &gojaScriptHost{}
}

type gojaScriptHost struct{}

type wrapper interface {
	constructor(call goja.ConstructorCall, r *goja.Runtime) *goja.Object
	storeInternal(value any, this *goja.Object)
}

type createWrapper func(instance *GojaContext) wrapper

type wrapperPrototypeInitializer interface {
	initializePrototype(prototype *goja.Object, r *goja.Runtime)
}

type class struct {
	name           string
	superClassName string
	wrapper        createWrapper
}

type classMap map[string]class

var globals classMap = make(classMap)

type function struct {
	Constructor *goja.Object
	Prototype   *goja.Object
	Wrapper     wrapper
}

type instanceInitializer interface {
	initObject(*goja.Object)
}

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
	vm := goja.New()
	vm.SetFieldNameMapper(propertyNameMapper{})
	result := &GojaContext{
		vm:           vm,
		clock:        clock.New(),
		window:       window,
		wrappedGoObj: goja.NewSymbol(internal_symbol_name),
		cachedNodes:  make(map[int32]goja.Value),
		classes:      make(map[string]*gojaClass),
	}
	for _, i := range factory.initializers {
		i.Configure(result)
	}

	globalThis := vm.GlobalObject()
	globalThis.DefineDataPropertySymbol(
		result.wrappedGoObj,
		vm.ToValue(window),
		goja.FLAG_FALSE,
		goja.FLAG_FALSE,
		goja.FLAG_FALSE,
	)
	globalThis.Set("window", globalThis)
	newEventLoopWrapper(result).initializeWindows(globalThis, vm)
	location := result.createLocationInstance()
	globalThis.DefineAccessorProperty("location", vm.ToValue(func(c *goja.FunctionCall) goja.Value {
		return location
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	globalThis.SetPrototype(result.classes["Window"].prototype)

	return result
}

func (d *gojaScriptHost) Close() {}
