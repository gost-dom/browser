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
	p := m.record.Evaluate(m.ctx.vm)

	if p.State() != sobek.PromiseStateFulfilled {
		return p.Result().Export().(error)
	}
	return m.ctx.clock.Tick()
}
