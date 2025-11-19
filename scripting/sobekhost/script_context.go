package sobekhost

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

type scriptContext struct {
	host         *gojaScriptHost
	vm           *sobek.Runtime
	clock        *clock.Clock
	window       html.Window
	classes      map[string]*class
	wrappedGoObj *sobek.Symbol
	cachedNodes  map[int32]sobek.Value
}

func (c *scriptContext) Clock() html.Clock { return c.clock }

func (i *scriptContext) Close() {}

func (i *scriptContext) logger() *slog.Logger {
	if i.window == nil {
		return nil
	}
	return i.window.Logger()
}

func (i *scriptContext) run(str string) (sobek.Value, error) {
	res, err := i.vm.RunString(str)
	i.clock.Tick()
	return res, err
}

func (i *scriptContext) Run(str string) error {
	s, e1 := i.Compile(str)
	return errors.Join(e1, s.Run())
}

func (i *scriptContext) Eval(str string) (res any, err error) {
	s, e1 := i.Compile(str)
	r, e2 := s.Eval()
	return r, errors.Join(e1, e2)
}

func (i *scriptContext) EvalCore(str string) (res any, err error) {
	return i.vm.RunString(str)
}

func (c *scriptContext) storeInternal(value any, obj *sobek.Object) {
	obj.DefineDataPropertySymbol(
		c.wrappedGoObj,
		c.vm.ToValue(value),
		sobek.FLAG_FALSE,
		sobek.FLAG_FALSE,
		sobek.FLAG_FALSE,
	)
	if e, ok := value.(entity.ObjectIder); ok {
		c.cachedNodes[e.ObjectId()] = obj
	}
}

func (i *scriptContext) RunFunction(str string, arguments ...any) (res any, err error) {
	var f sobek.Value
	if f, err = i.vm.RunString(str); err == nil {
		if c, ok := sobek.AssertFunction(f); !ok {
			err = errors.New("GojaContext.RunFunction: script is not a function")
		} else {
			values := make([]sobek.Value, len(arguments))
			for i, a := range arguments {
				var ok bool
				if values[i], ok = a.(sobek.Value); !ok {
					err = fmt.Errorf("GojaContext.RunFunction: argument %d was not a goja Value", i)
				}
			}
			res, err = c(sobek.Undefined(), values...)
		}
	}
	return
}

// Export create a native Go value out of a javascript value. The value argument
// must be a [sobek.Value] instance.
//
// This function is intended to be used only for test purposes. The value has an
// [any] type as the tests are not supposed to know the details of the
// underlying engine.
//
// The value is expected to be the ourput of [RunFunction] or [EvalCore]
//
// An error will be returned if the value is not a goja Value, or the value
// could not be converted to a native Go object
func (i *scriptContext) Export(value any) (res any, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("GojaContext.Export: %v", r)
		}
	}()
	if gv, ok := value.(sobek.Value); ok {
		res = gv.Export()
	} else {
		err = fmt.Errorf("GojaContext.Export: Value not a goja value: %v", value)
	}
	return
}

func (m *scriptContext) createLocationInstance() *sobek.Object {
	location, err := m.classes["Location"].NewInstance(m.window.Location())
	if err != nil {
		panic(err)
	}
	return location.(gojaObject).obj
}

func (c *scriptContext) CreateFunction(name string, cb js.FunctionCallback[jsTypeParam]) {
	c.vm.Set(name, wrapJSCallback(c, cb))
}

func (c *scriptContext) RunScript(script, src string) {
	_, err := c.vm.RunScript(src, script)
	if err != nil {
		fmt.Println("RUN SCRIPT FAIL", script, src)
		panic(err)
	}
}

func (c *scriptContext) CreateClass(
	name string, extends js.Class[jsTypeParam],
	cb js.FunctionCallback[jsTypeParam],
) js.Class[jsTypeParam] {
	cls := &class{ctx: c, cb: cb, instanceAttrs: make(map[string]attributeHandler)}
	constructor := c.vm.ToValue(cls.constructorCb).(*sobek.Object)
	constructor.DefineDataProperty(
		"name",
		c.vm.ToValue(name),
		sobek.FLAG_NOT_SET,
		sobek.FLAG_NOT_SET,
		sobek.FLAG_NOT_SET,
	)
	cls.prototype = constructor.Get("prototype").(*sobek.Object)
	c.vm.Set(name, constructor)
	c.classes[name] = cls

	if extends != nil {
		if superclass, ok := extends.(*class); ok {
			cls.prototype.SetPrototype(superclass.prototype)
		} else {
			panic(fmt.Sprintf("Superclass not installed for %s. extends: %+v", name, extends))
		}
	}

	return cls
}

func (class *class) constructorCb(call sobek.ConstructorCall, r *sobek.Runtime) *sobek.Object {
	class.installInstance(&call.This, nil)
	class.cb(newGojaCallbackContext(class.ctx, call))
	return nil
}

func (class *class) installInstance(this **sobek.Object, native any) {
	for _, v := range class.instanceAttrs {
		v.install(*this)
	}

	if class.namedHandlerCallbacks != nil {
		// This implementation is somewhat fragile if the object need own
		// properties. See comment below.
		obj := *this
		proto := *this
		*this = class.ctx.vm.NewDynamicObject(&gojaDynamicObject{
			ctx:   class.ctx,
			cbs:   *class.namedHandlerCallbacks,
			this:  obj,
			scope: newCallbackScope(class.ctx, proto, native),
		})
		(*this).SetPrototype(class.prototype)
	}
	if class.indexedHandler != nil {
		// TODO: Fix prototype for indexed property handlers. Due to lack of
		// support for internal values in goja, and because a "Dynamic Object"
		// cannot have own symbol properties, an artificial prototype is
		// inserted between the instance and the correct prototype, in order to
		// be able to retrieve the internal instance.
		proto := *this
		*this = class.ctx.vm.NewDynamicArray(&gojaDynamicArray{
			ctx:   class.ctx,
			cbs:   *class.indexedHandler,
			this:  *this,
			scope: newCallbackScope(class.ctx, proto, native),
		})
		(*this).SetPrototype(proto)
	}
}

type attributeHandler struct {
	ctx    *scriptContext
	name   string
	getter js.FunctionCallback[jsTypeParam]
	setter js.FunctionCallback[jsTypeParam]
}

func (h attributeHandler) install(object *sobek.Object) {
	var g, s sobek.Value
	if h.getter != nil {
		g = wrapJSCallback(h.ctx, h.getter)
	}
	if h.setter != nil {
		s = wrapJSCallback(h.ctx, h.setter)
	}
	object.DefineAccessorProperty(h.name, g, s,
		sobek.FLAG_TRUE, // configurable
		sobek.FLAG_TRUE, // enumerable
	)
}

func newGojaCallbackContext(
	ctx *scriptContext,
	call sobek.ConstructorCall,
) *callbackContext {
	return &callbackContext{
		newCallbackScope(ctx, call.This, nil),
		call.Arguments, 0,
	}
}

func (c *scriptContext) Compile(script string) (html.Script, error) {
	return GojaScript{c, script}, nil
}

func (c *scriptContext) DownloadScript(script string) (html.Script, error) {
	return nil, errors.New("TODO")
}

func (c *scriptContext) DownloadModule(script string) (result html.Script, err error) {
	resolver := sobekResolver{
		c.host,
		c,
		make(map[sobek.ModuleRecord]string),
		make(map[string]sobek.ModuleRecord),
	}
	rec, err := resolver.resolveModule(c.window.LocationHREF(), script)
	if err == nil {
		result = sobekModule{c, rec}
	}
	return
	// return nil, errors.New("gojahost: ECMAScript modules not supported by gojehost")
}

/* -------- GojaScript -------- */

type GojaScript struct {
	context *scriptContext
	script  string
}

func (s GojaScript) Run() error {
	_, err := s.context.run(s.script)
	return err
}

func (s GojaScript) Eval() (res any, err error) {
	if gojaVal, err := s.context.run(s.script); err == nil {
		if sobek.IsNull(gojaVal) || sobek.IsUndefined(gojaVal) {
			return nil, nil
		}
		if obj := gojaVal.ToObject(s.context.vm); obj != nil {
			if v := obj.GetSymbol(s.context.wrappedGoObj); v != nil {
				if e := v.Export(); e != nil {
					return e, nil
				}
			}
		}
		return gojaVal.Export(), nil
	} else {
		return nil, err
	}
}
