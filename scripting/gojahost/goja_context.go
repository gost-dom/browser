package gojahost

import (
	"errors"
	"fmt"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"

	"github.com/dop251/goja"
)

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
	s, e1 := i.Compile(str)
	return errors.Join(e1, s.Run())
}

func (i *GojaContext) Eval(str string) (res any, err error) {
	s, e1 := i.Compile(str)
	r, e2 := s.Eval()
	return r, errors.Join(e1, e2)
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

func (c *GojaContext) Compile(script string) (html.Script, error) {
	return GojaScript{c, script}, nil
}

func (c *GojaContext) DownloadScript(script string) (html.Script, error) {
	return nil, errors.New("TODO")
}

func (c *GojaContext) DownloadModule(script string) (html.Script, error) {
	return nil, errors.New("gojahost: ECMAScript modules not supported by gojehost")
}

/* -------- GojaScript -------- */

type GojaScript struct {
	context *GojaContext
	script  string
}

func (s GojaScript) Run() error {
	_, err := s.context.run(s.script)
	return err
}

func (s GojaScript) Eval() (res any, err error) {
	if gojaVal, err := s.context.run(s.script); err == nil {
		return gojaVal.Export(), nil
	} else {
		return nil, err
	}
}
