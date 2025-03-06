package entity

import "sync/atomic"

// An ObjectId uniquely identifies an element in the DOM. It is meant for
// internal use only, and shouldn't be used by users of the library.
//
// The value is a 32bit integer so it can accurately be represented by a
// JavaScript number.
type ObjectId = int32

var idSeq atomic.Int32

// NewObjectId returns a new guaranteed atomically unique ObjectId.
//
// Warning: This solution is temporary, and a different solution is intended to
// be used, so function will likely disappear in the future.
func NewObjectId() ObjectId {
	return idSeq.Add(1)
}

// An ObjectIder provides a unique identifier of an object that may be retrieved
// from the DOM. It is part of a solution to ensure the same JS object is
// returned for the same DOM element.
//
// Warning: This solution is temporary, and a different solution is intended to
// be used. Do not rely on this value.
type ObjectIder interface {
	ObjectId() ObjectId
}

// Entity is the default Entity implementation. The zero value will generate a
// unique [ObjectId] the first time it is read.
type Entity struct {
	objectId ObjectId
}

func (b *Entity) ObjectId() ObjectId {
	if b.objectId == 0 {
		b.objectId = NewObjectId()
	}
	return b.objectId
}
