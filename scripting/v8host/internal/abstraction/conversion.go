package abstraction

import "fmt"

func As[T any](val any, err error) (rtnVal T, rtnErr error) {
	if err != nil {
		rtnErr = err
		return
	}
	var ok bool
	if rtnVal, ok = val.(T); !ok {
		rtnErr = fmt.Errorf("value %+v is not assignable to requested type %T", val, rtnVal)
	}
	return
}
