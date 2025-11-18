package sobekhost

import (
	"errors"

	"github.com/grafana/sobek"
)

type sobekModule struct {
	vm *sobek.Runtime
}

func (m sobekModule) Eval() (any, error) {
	return nil, errors.New("Not implemented")
}
func (m sobekModule) Run() error {
	return errors.New("sobekhost.Run not implemented")
}
