package Registry

import (
	"github.com/yindaheng98/go-utility/Emitter"
	"gogistery/util/emitters"
)

type events struct {
	NewConnection     *emitters.RegistrantInfoEmitter
	UpdateConnection  *emitters.RegistrantInfoEmitter
	ConnectionTimeout *emitters.RegistrantInfoEmitter
	Disconnection     *emitters.RegistrantInfoEmitter
	Error             *Emitter.ErrorEmitter
}

func newEvents() *events {
	return &events{
		emitters.NewRegistrantInfoEmitter(),
		emitters.NewRegistrantInfoEmitter(),
		emitters.NewRegistrantInfoEmitter(),
		emitters.NewRegistrantInfoEmitter(),
		Emitter.NewErrorEmitter()}
}

func (e *events) EnableAll() {
	e.NewConnection.Enable()
	e.UpdateConnection.Enable()
	e.ConnectionTimeout.Enable()
	e.Disconnection.Enable()
	e.Error.Enable()
}

func (e *events) DisableAll() {
	e.NewConnection.Disable()
	e.UpdateConnection.Disable()
	e.ConnectionTimeout.Disable()
	e.Disconnection.Disable()
	e.Error.Enable()
}
