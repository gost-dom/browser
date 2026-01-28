package sobekengine

import (
	"errors"
	"fmt"

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
			res := p.Result().Export()
			err, ok := res.(error)
			if !ok {
				err = fmt.Errorf("run module: %v", res)
			}
			return err
		}
		return nil
	})
}
