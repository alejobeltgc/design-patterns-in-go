package concretecommands

import "designpatterns/behavioral/command/remote/devices"

type GarageDoorOpenCommand struct {
	GarageDoor *devices.GarageDoor
}

func NewGarageDoorOpenCommand(garageDoor *devices.GarageDoor) *GarageDoorOpenCommand {
	return &GarageDoorOpenCommand{
		GarageDoor: garageDoor,
	}
}

func (l *GarageDoorOpenCommand) Execute() {
	l.GarageDoor.Up()
}

func (l *GarageDoorOpenCommand) Undo() {
	l.GarageDoor.Down()
}
