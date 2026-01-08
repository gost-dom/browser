package entity

import (
	"reflect"
	"sync/atomic"
)

// An ObjectId uniquely identifies an element in the DOM. It is meant for
// internal use only, and shouldn't be used by users of the library.
//
// The value is a 32bit integer so it can accurately be represented by a
// JavaScript number.
//
// Deprecated: This will likely be moved to an internal package
type ObjectId = int32

var idSeq atomic.Int32

// NewObjectId returns a new guaranteed atomically unique ObjectId.
//
// Deprecated: This will likely be moved to an internal package
func NewObjectId() ObjectId {
	return idSeq.Add(1)
}

// An ObjectIder provides a unique identifier of an object that may be retrieved
// from the DOM. It is part of a solution to ensure the same JS object is
// returned for the same DOM element.
//
// Deprecated: This will likely be moved to an internal package
type ObjectIder interface {
	ObjectId() ObjectId
}

// Components allow storing and retrieving components associated with an entity.
// This is very inspired by the idea of an Entity Component System.
//
// Keys must be non-nil and comparable. Methods will panic if the key is nil or
// not comparable. When using string values as keys, it's advisable to create a
// new type for the keys to avoid name collision with keys from other parts of
// the code.
//
// Example usages:
//
//   - The JavaScript wrapper for a component is stored in the component.
//   - Test code stores a *testing.T in the window object, allowing assertions
//     in JavaScript to integrate with Go tests.
//
// See also: https://en.wikipedia.org/wiki/Entity_component_system
type Components interface {
	// Methods are kept unexported to allow refactoring
	component(key any) (val any, found bool)
	setComponent(key any, val any)
}

type componentEntry struct {
	key reflect.Value
	val any
}

// Entity is the default Components implementation.
//
// Deprecated: This will likely be renamed or moved to an internal package
type Entity struct {
	components []componentEntry
}

func parseKey(key any) reflect.Value {
	if key == nil {
		panic("nil key")
	}
	v := reflect.ValueOf(key)
	if !v.Comparable() {
		panic("key is not comparable")
	}
	return v
}

func (e *Entity) component(key any) (res any, ok bool) {
	res, _, ok = e.find(parseKey(key))
	return
}

func (e *Entity) setComponent(key any, val any) {
	v := parseKey(key)
	_, idx, ok := e.find(v)
	if ok {
		e.components[idx].val = val
	} else {
		e.components = append(e.components, componentEntry{v, val})
	}
}

func Component[T any](e Components, key any) (res T, ok bool) {
	var val any
	if val, ok = e.component(key); !ok {
		return
	}
	res, ok = val.(T)
	return
}

func SetComponent(e Components, key any, val any) { e.setComponent(key, val) }

func ComponentType[T any](e Components) (T, bool) {
	return Component[T](e, reflect.TypeFor[T]())
}

func SetComponentType[T any](e Components, val T) { e.setComponent(reflect.TypeFor[T](), val) }

func (e *Entity) find(v reflect.Value) (any, int, bool) {
	for i, pair := range e.components {
		if pair.key.Equal(v) {
			return pair.val, i, true
		}
	}
	return nil, -1, false
}
