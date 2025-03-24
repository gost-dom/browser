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

func (s *setups) tryAppend(val any) {
	if setup, ok := val.(Setuper); ok {
		s.append(setup)
	}
}

func (s *NullSetup) Setup() {}

type Fixture struct {
	testing.TB
}

func (f *Fixture) SetTB(tb testing.TB) { f.TB = tb }

func (f *FixtureSetup[T]) Setup() Setuper {
	vType := reflect.TypeFor[T]()
	f.TB.Helper()
	if vType.Kind() != reflect.Pointer {
		f.TB.Fatalf("InitFixture: Fixture must be a pointer. Actual type: %s", vType.Name())
	}
	vType = vType.Elem()
	f.pkgPath = vType.PkgPath()
	f.depVals = make([]reflect.Value, len(f.Deps))
	for i, d := range f.Deps {
		f.depVals[i] = reflect.ValueOf(d)
	}
	return f.setup(reflect.ValueOf(f.Fixture))
}

func (f FixtureSetup[T]) defaultInclude(val reflect.Value) bool {
	var underlyingType = val.Type()
	if underlyingType.Kind() == reflect.Pointer {
		underlyingType = underlyingType.Elem()
	}
	return strings.HasSuffix(underlyingType.Name(), "Fixture")
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

	depVals []reflect.Value
	pkgPath string
}

func (f *FixtureSetup[T]) setupStruct(val reflect.Value, typ reflect.Type) *setups {
	var setups = new(setups)
fields:
	for _, field := range reflect.VisibleFields(typ) {
		if len(field.Index) > 1 || !field.IsExported() {
			// Don't set fields of embedded or unexported fields
			continue
		}
		fieldVal := val.FieldByIndex(field.Index)
		if !f.include(fieldVal) {
			continue
		}
		for _, depVal := range f.depVals {
			if field.Type == depVal.Type() {
				fieldVal.Set(depVal)
				continue fields
			}
		}
		if field.Type.Kind() == reflect.Pointer && fieldVal.IsNil() {
			fieldVal.Set(reflect.New(field.Type.Elem()))
			f.depVals = append(f.depVals, fieldVal)
		}
		setups.append(f.setup(fieldVal))
	}
	return setups
}

func (f *FixtureSetup[T]) setup(val reflect.Value) Setuper {
	var setups = new(setups)
	if val.Kind() == reflect.Pointer {
		if val.IsNil() {
			return &NullSetup{}
		}
		val = val.Elem()
	}

	typ := val.Type()
	if typ.Kind() == reflect.Struct {
		setups.append(f.setupStruct(val, typ))
	}

	if !val.CanAddr() {
		asAny := val.Interface()
		setups.tryAppend(asAny)
		f.tryInit(asAny)
	} else {
		// val must be addressable, as both Setup and Init are mutating functions.
		//
		// Val itself may be a non-pointer field in a pointer struct, which
		// means val.Interface() itself is not a Setuper or Initer*, but we can
		// still get a pointer using Addr() because it's inside an addressable
		// struct.
		//
		// \* While a non-pointer val may _implement_ the interfaces, the
		// implementation would be wrong, as they couldn't mutate.

		asAny := val.Addr().Interface()
		setups.tryAppend(asAny)
		f.tryInit(asAny)
	}
	return setups
}

func (s *FixtureSetup[T]) tryInit(val any) {
	if init, ok := val.(FixtureInit); ok {
		init.SetTB(s.TB)
	}
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
