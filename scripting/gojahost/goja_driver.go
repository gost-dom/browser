package gojahost

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/scripting"

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
	installClass("EventTarget", "", newEventTargetWrapper)
	installClass("Window", "Node", newWindowWrapper)
	installClass("Document", "Node", newDocumentWrapper)
	installClass("HTMLDocument", "Document", newHTMLDocumentWrapper)
	installClass("Event", "", newEventWrapperAsWrapper)
	installClass("CustomEvent", "Event", newCustomEventWrapper)
	installClass("Element", "Node", newElementWrapper)
	installClass("HTMLElement", "Element", newGenericElementWrapper)

	for _, cls := range scripting.HtmlElements {
		if _, found := globals[cls]; !found {
			installClass(cls, "HTMLElement", newGenericElementWrapper)
		}
	}
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

func (d *GojaContext) installGlobals(classes classMap) {
	d.globals = make(map[string]function)
	var assertGlobal func(class) function
	assertGlobal = func(class class) function {
		name := class.name
		wrapper := class.wrapper(d)
		if constructor, alreadyInstalled := d.globals[name]; alreadyInstalled {
			return constructor
		}
		constructor := d.vm.ToValue(wrapper.constructor).(*goja.Object)
		constructor.DefineDataProperty(
			"name",
			d.vm.ToValue(name),
			goja.FLAG_NOT_SET,
			goja.FLAG_NOT_SET,
			goja.FLAG_NOT_SET,
		)
		prototype := constructor.Get("prototype").(*goja.Object)
		result := function{constructor, prototype, wrapper}
		d.vm.Set(name, constructor)
		d.globals[name] = result

		if super := class.superClassName; super != "" {
			if superclass, found := classes[super]; found {
				superPrototype := assertGlobal(superclass).Prototype
				prototype.SetPrototype(superPrototype)
			} else {
				panic(fmt.Sprintf("Superclass not installed for %s. Superclass: %s", name, super))
			}
		}

		if initializer, ok := wrapper.(wrapperPrototypeInitializer); ok {
			initializer.initializePrototype(prototype, d.vm)
		}

		return result
	}
	for _, class := range classes {
		assertGlobal(class)
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
	}
	result.installGlobals(globals)

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
	globalThis.SetPrototype(result.globals["Window"].Prototype)

	return result
}

func (m *GojaContext) createLocationInstance() *goja.Object {
	location := m.vm.CreateObject(m.globals["Location"].Prototype)
	location.DefineDataPropertySymbol(
		m.wrappedGoObj,
		m.vm.ToValue(m.window.Location()),
		goja.FLAG_FALSE,
		goja.FLAG_FALSE,
		goja.FLAG_FALSE,
	)
	return location
}
func (d *gojaScriptHost) Close() {}

type GojaContext struct {
	vm           *goja.Runtime
	clock        *clock.Clock
	window       html.Window
	globals      map[string]function
	wrappedGoObj *goja.Symbol
	cachedNodes  map[int32]goja.Value
}

func (c *GojaContext) Clock() html.Clock { return c.clock }

func (i *GojaContext) Close() {}

func (i *GojaContext) run(str string) (goja.Value, error) {
	res, err := i.vm.RunString(str)
	i.clock.Tick()
	return res, err
}

func (i *GojaContext) Run(str string) error {
	_, err := i.run(str)
	return err
}

func (i *GojaContext) Eval(str string) (res any, err error) {
	if gojaVal, err := i.run(str); err == nil {
		return gojaVal.Export(), nil
	} else {
		return nil, err
	}
}

func (i *GojaContext) EvalCore(str string) (res any, err error) {
	return i.vm.RunString(str)
}

func (i *GojaContext) RunFunction(str string, arguments ...any) (res any, err error) {
	var f goja.Value
	if f, err = i.vm.RunString(str); err == nil {
		if c, ok := goja.AssertFunction(f); !ok {
			err = errors.New("GojaContext.RunFunction: script is not a function")
		} else {
			values := make([]goja.Value, len(arguments))
			for i, a := range arguments {
				var ok bool
				if values[i], ok = a.(goja.Value); !ok {
					err = fmt.Errorf("GojaContext.RunFunction: argument %d was not a goja Value", i)
				}
			}
			res, err = c(goja.Undefined(), values...)
		}
	}
	return
}

// Export create a native Go value out of a javascript value. The value argument
// must be a [goja.Value] instance.
//
// This function is intended to be used only for test purposes. The value has an
// [any] type as the tests are not supposed to know the details of the
// underlying engine.
//
// The value is expected to be the ourput of [RunFunction] or [EvalCore]
//
// An error will be returned if the value is not a goja Value, or the value
// could not be converted to a native Go object
func (i *GojaContext) Export(value any) (res any, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("GojaContext.Export: %v", r)
		}
	}()
	if gv, ok := value.(goja.Value); ok {
		res = gv.Export()
	} else {
		err = fmt.Errorf("GojaContext.Export: Value not a goja value: %v", value)
	}
	return
}
