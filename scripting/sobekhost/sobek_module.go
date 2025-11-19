package sobekhost

import (
	"errors"

	"github.com/grafana/sobek"
)

type sobekModule struct {
	ctx    *GojaContext
	record sobek.ModuleRecord
}

func (m sobekModule) Eval() (any, error) {
	return nil, errors.New("Not implemented")
}

func (m sobekModule) Run() error {
	m.ctx.logger().Debug("Evaluate module", "vm", m.ctx.vm)
	// TODO: Handle promise return value

	// if err := m.record.Link(); err != nil {
	// 	return err
	// }
	m.record.Evaluate(m.ctx.vm)
	return nil
}
