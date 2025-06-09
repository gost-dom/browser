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

func installClass(name string, superClassName string, wrapper createWrapper) {
	if _, found := globals[name]; found {
		panic("Class already installed")
	}
	globals[name] = class{name, superClassName, wrapper}
}

func init() {
	// installClass("EventTarget", "", newEventTargetWrapper)
	// installClass("Window", "Node", newWindowWrapper)
	// installClass("Document", "Node", newDocumentWrapper)
	// installClass("HTMLDocument", "Document", newHTMLDocumentWrapper)
	// installClass("CustomEvent", "Event", newCustomEventWrapper)
	// installClass("Element", "Node", newElementWrapper)
	// installClass("HTMLElement", "Element", newGenericElementWrapper)
	//
	// for _, cls := range codec.HtmlElements {
	// 	if _, found := globals[cls]; !found {
	// 		installClass(cls, "HTMLElement", newGenericElementWrapper)
	// 	}
	// }
}

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

// func (d *GojaContext) installGlobals(classes classMap) {
// 	d.globals = make(map[string]function)
// 	d.classes = make(map[string]*gojaClass)
// 	var assertGlobal func(class) function
// 	assertGlobal = func(class class) function {
// 		name := class.name
// 		wrapper := class.wrapper(d)
// 		if constructor, alreadyInstalled := d.globals[name]; alreadyInstalled {
// 			return constructor
// 		}
// 		constructor := d.vm.ToValue(wrapper.constructor).(*goja.Object)
// 		constructor.DefineDataProperty(
// 			"name",
// 			d.vm.ToValue(name),
// 			goja.FLAG_NOT_SET,
// 			goja.FLAG_NOT_SET,
// 			goja.FLAG_NOT_SET,
// 		)
// 		prototype := constructor.Get("prototype").(*goja.Object)
// 		result := function{constructor, prototype, wrapper}
// 		d.vm.Set(name, constructor)
// 		d.globals[name] = result
//
// 		if super := class.superClassName; super != "" {
// 			if superclass, found := classes[super]; found {
// 				superPrototype := assertGlobal(superclass).Prototype
// 				prototype.SetPrototype(superPrototype)
// 			} else {
// 				panic(fmt.Sprintf("Superclass not installed for %s. Superclass: %s", name, super))
// 			}
// 		}
//
// 		if initializer, ok := wrapper.(wrapperPrototypeInitializer); ok {
// 			initializer.initializePrototype(prototype, d.vm)
// 		}
//
// 		return result
// 	}
// 	for _, class := range classes {
// 		assertGlobal(class)
// 	}
// }

func (d *gojaScriptHost) NewContext(window html.Window) html.ScriptContext {
	vm := goja.New()
	vm.SetFieldNameMapper(propertyNameMapper{})
	result := &GojaContext{
		vm:           vm,
		clock:        clock.New(),
		window:       window,
		wrappedGoObj: goja.NewSymbol(internal_symbol_name),
		cachedNodes:  make(map[int32]goja.Value),
		// globals:      make(map[string]function),
		classes: make(map[string]*gojaClass),
	}
	for _, i := range factory.initializers {
		i.Configure(result)
	}
	// factory.result.installGlobals(globals)

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
	globalThis.DefineAccessorProperty("document", vm.ToValue(func(c *goja.FunctionCall) goja.Value {
		return result.toNode(window.Document())
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	location := result.createLocationInstance()
	globalThis.DefineAccessorProperty("location", vm.ToValue(func(c *goja.FunctionCall) goja.Value {
		return location
	}), nil, goja.FLAG_FALSE, goja.FLAG_TRUE)
	globalThis.SetPrototype(result.classes["Window"].prototype)

	return result
}

func (d *gojaScriptHost) Close() {}
