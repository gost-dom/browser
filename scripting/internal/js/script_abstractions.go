package js

import (
	"fmt"
	"reflect"
)

func As[T any](val any, err error) (rtnVal T, rtnErr error) {
	if err != nil {
		rtnErr = err
		return
	}
	var ok bool
	if rtnVal, ok = val.(T); !ok {
		rtnErr = fmt.Errorf(
			"value %+v (%T) is not assignable to requested type %v",
			val,
			val,
			reflect.TypeFor[T](),
		)
	}
	return
}
