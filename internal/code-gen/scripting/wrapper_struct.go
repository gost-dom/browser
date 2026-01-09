package scripting

import (
	"fmt"

	"github.com/gost-dom/code-gen/scripting/model"
	g "github.com/gost-dom/generators"
)

type WrapperStruct struct {
	Data model.ESConstructorData
}

func (ge WrapperStruct) InitializerName() string {
	return fmt.Sprintf("Initialize%s", ge.IdlName())
}

func (ge WrapperStruct) SpecName() string { return ge.Data.SpecName() }

func (ge WrapperStruct) IdlName() string { return ge.Data.Name() }

func Initializer(d model.ESConstructorData) g.Value {
	ws := WrapperStruct{d}
	return g.NewValue(ws.InitializerName())
}

func (ge WrapperStruct) Callbacks() CallbackMethods {
	return CallbackMethods{ge}
}
