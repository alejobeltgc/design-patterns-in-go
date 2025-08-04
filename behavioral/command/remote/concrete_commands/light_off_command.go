package concretecommands

import "designpatterns/behavioral/command/remote/devices"

type LightOffCommand struct {
	Light *devices.Light
}

func NewLightOffCommand(light *devices.Light) *LightOffCommand {
	return &LightOffCommand{
		Light: light,
	}
}

func (l *LightOffCommand) Execute() {
	l.Light.Off()
}

func (l *LightOffCommand) Undo() {
	l.Light.On()
}
