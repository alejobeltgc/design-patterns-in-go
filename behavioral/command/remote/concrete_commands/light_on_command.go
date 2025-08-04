package concretecommands

import "designpatterns/behavioral/command/remote/devices"

type LightOnCommand struct {
	Light *devices.Light
}

func NewLightOnCommand(light *devices.Light) *LightOnCommand {
	return &LightOnCommand{
		Light: light,
	}
}

func (l *LightOnCommand) Execute() {
	l.Light.On()
}

func (l *LightOnCommand) Undo() {
	l.Light.Off()
}
