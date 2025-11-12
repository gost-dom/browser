package js

import "fmt"

type ConstructorFactory[T any] = func(ScriptEngine[T], Class[T]) Class[T]

type classSpec[T any] struct {
	name           string
	superClassName string
	factory        ConstructorFactory[T]
}

// ClassBuilder provides a simplified way of registering classes in global
// scope. Method [CreateGlobals] initializes a script engine with the classes
// registered using [Register]. This handles the correct order of registration,
// creating superclasses before subclasses.
type ClassBuilder[T any] struct {
	classes map[string]classSpec[T]
}

func NewClassBuilder[T any]() ClassBuilder[T] {
	return ClassBuilder[T]{make(map[string]classSpec[T])}
}
func (b ClassBuilder[T]) HasClass(name string) bool {
	_, ok := b.classes[name]
	return ok
}

func (r ClassBuilder[T]) Register(
	className, superClassName string,
	fact InitializerFactory[T, Initializer[T]],
) {
	spec := classSpec[T]{
		className, superClassName, func(host ScriptEngine[T], extends Class[T]) Class[T] {
			wrapper := fact(host)
			res := host.CreateClass(className, extends, wrapper.Constructor)
			wrapper.Initialize(res)
			return res
		},
	}
	if _, ok := r.classes[className]; ok {
		panic("Same class added twice: " + className)
	}
	if superClassName == "" {
		r.classes[className] = spec
		return
	}
	parent, parentFound := r.classes[superClassName]
	for parentFound {
		if parent.superClassName == className {
			panic("Recursive class parents" + className)
		}
		parent, parentFound = r.classes[parent.superClassName]
	}
	r.classes[className] = spec
}

type Initializer[T any] interface {
	Constructor(CallbackContext[T]) (Value[T], error)
	Initialize(Class[T])
}

type InitializerFactory[T any, U Initializer[T]] = func(ScriptEngine[T]) U

// CreateGlobals registers classes in global scope, as well as the prototype and
// instance operations and attributes. The function conforms to the
// [ConfigurerFunc] type.
// The function creates the classes in the correct order, creating superclasses
// before subclasses.
func (c *ClassBuilder[T]) CreateGlobals(host ScriptEngine[T]) {
	var iter func(class classSpec[T]) Class[T]
	uniqueNames := make(map[string]Class[T])
	iter = func(class classSpec[T]) Class[T] {
		if constructor, found := uniqueNames[class.name]; found {
			return constructor
		}
		var superClassConstructor Class[T]
		if class.superClassName != "" {
			superClassSpec, found := c.classes[class.superClassName]
			if !found {
				panic(fmt.Sprintf(
					"Missing super class spec. Class: %s. Super: %s",
					class.name, class.superClassName,
				))
			}
			superClassConstructor = iter(superClassSpec)
		}
		constructor := class.factory(host, superClassConstructor)
		uniqueNames[class.name] = constructor
		return constructor
	}
	for _, class := range c.classes {
		iter(class)
	}
}

func RegisterClass[T any, U Initializer[T], V InitializerFactory[T, U]](
	reg ClassBuilder[T],
	className, superClassName string,
	constructorFactory V) {
	reg.Register(
		className,
		superClassName,
		func(h ScriptEngine[T]) Initializer[T] { return constructorFactory(h) },
	)
}
