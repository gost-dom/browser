package fixture

import (
	"reflect"
	"strings"
	"testing"
)

type FixtureInit interface {
	SetTB(testing.TB)
}

type Setuper interface{ Setup() }

// NullSetuper is just a Setuper that ignores setup calls on the null instance
type NullSetup struct{}

type setups []Setuper

func (s setups) Setup() {
	for _, ss := range s {
		ss.Setup()
	}
}

func (s *setups) append(setup Setuper) {
	*s = append(*s, setup)
}

func (s *NullSetup) Setup() {}

type Fixture struct {
	testing.TB
}

func (f *Fixture) SetTB(tb testing.TB) { f.TB = tb }

func typePkgPath(t reflect.Type) (string, bool) {
	k := t.Kind()
	if k == reflect.Pointer {
		t = t.Elem()
		k = t.Kind()
	}
	if k == reflect.Struct {
		return t.PkgPath(), true
	}
	return "", false
}

// valuePkgPath is specifically for filtering struct types based on their
// package path. The use of an ok return value is solely for this specific
// purpose alone; and the function may not be a good fit in other applications.
func valuePkgPath(val reflect.Value) (string, bool) {
	t := val.Type()
	return typePkgPath(t)
}

func (f *FixtureSetup[T]) Setup() Setuper {
	vType := reflect.TypeFor[T]()
	f.TB.Helper()
	if vType.Kind() != reflect.Pointer {
		f.TB.Fatalf("InitFixture: Fixture must be a pointer. Actual type: %s", vType.Name())
	}
	vType = vType.Elem()
	f.pkgPath = vType.PkgPath()
	return f.setup(reflect.ValueOf(f.Fixture))
}

func (f FixtureSetup[T]) defaultInclude(val reflect.Value) bool {
	var currPath, expPath string
	var ok bool
	if currPath, ok = valuePkgPath(val); ok {
		expPath, ok = typePkgPath(reflect.TypeFor[T]())
	}
	if !ok {
		return false
	}
	return strings.HasPrefix(currPath, expPath)
}

func (f FixtureSetup[T]) include(val reflect.Value) bool {
	if f.Include != nil {
		return f.Include(val)
	}
	return f.defaultInclude(val)
}

type FixtureSetup[T FixtureInit] struct {
	TB      testing.TB
	Fixture T
	Include func(val reflect.Value) bool
	Deps    []any

	pkgPath string
}

func (f FixtureSetup[T]) setup(val reflect.Value) Setuper {
	var setups = new(setups)
	if !f.include(val) {
		return setups
	}
	if val.Kind() == reflect.Pointer {
		if val.IsNil() {
			return &NullSetup{}
		}
		val = val.Elem()
	}
	typ := val.Type()
	for _, field := range reflect.VisibleFields(typ) {
		if len(field.Index) > 1 { // Don't set fields of embedded members
			continue
		}
		for _, dep := range f.Deps {
			depVal := reflect.ValueOf(dep)
			if field.Type == reflect.TypeOf(dep) {
				val.FieldByIndex(field.Index).Set(depVal)
				continue
			}
		}
		setups.append(f.setup(val.FieldByIndex(field.Index)))
	}
	actualVal := val.Interface()
	if setup, ok := actualVal.(Setuper); ok {
		setups.append(setup)
	} else if val.CanAddr() {
		if setup, ok := val.Addr().Interface().(Setuper); ok {
			setups.append(setup)
		}
	}
	if init, ok := actualVal.(FixtureInit); ok {
		init.SetTB(f.TB)
	} else if val.CanAddr() {
		if init, ok := val.Addr().Interface().(FixtureInit); ok {
			init.SetTB(f.TB)
		}
	}
	// var res *NullSetup
	return setups
}

func InitFixture[T FixtureInit](
	t testing.TB,
	fixture T,
	deps ...any,
) (T, Setuper) {
	f := &FixtureSetup[T]{
		TB:      t,
		Fixture: fixture,
		Deps:    deps,
	}
	return fixture, f.Setup()
}
