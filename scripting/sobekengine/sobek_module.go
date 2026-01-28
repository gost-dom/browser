package sobekengine

import (
	"errors"

	"github.com/grafana/sobek"
)

type sobekModule struct {
	ctx    *scriptContext
	record sobek.ModuleRecord
}

func (m sobekModule) Eval() (any, error) {
	return nil, errors.New("Not implemented")
}

func (m sobekModule) Run() error {
	return m.ctx.do(func() error {
		m.ctx.logger().Debug("Evaluate module", "vm", m.ctx.vm)
		p := m.record.Evaluate(m.ctx.vm)

		if p.State() != sobek.PromiseStateFulfilled {
			return p.Result().Export().(error)
		}
		return nil
	})
}
