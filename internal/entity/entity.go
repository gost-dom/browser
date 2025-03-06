package entity

import "sync/atomic"

// An ObjectId uniquely identifies an element in the DOM. It is meant for
// internal use only, and shouldn't be used by users of the library.
//
// The value is a 32bit integer so it can accurately be represented by a
// JavaScript number.
type ObjectId = int32

var idSeq atomic.Int32

func NewObjectId() ObjectId {
	return idSeq.Add(1)
}

// An Entity provides a unique identifier of an object that may be retrieved
// from the DOM. It is part of a solution to ensure the same JS object is
// returned for the same DOM element.
//
// Warning: This solution is temporary, and a different solution is intended to
// be used. Do not rely on this value.
type Entity interface {
	ObjectId() ObjectId
}

type Base struct {
	objectId ObjectId
}

func New() Entity {
	return Base{NewObjectId()}
}

func (b Base) ObjectId() ObjectId {
	if b.objectId == 0 {
		b.objectId = NewObjectId()
	}
	return b.objectId
}
