package concretecommands

import "designpatterns/behavioral/command/remote/devices"

type GarageDoorDownCommand struct {
	GarageDoor *devices.GarageDoor
}

func NewGarageDoorDownCommand(garageDoor *devices.GarageDoor) *GarageDoorDownCommand {
	return &GarageDoorDownCommand{
		GarageDoor: garageDoor,
	}
}

func (g *GarageDoorDownCommand) Execute() {
	g.GarageDoor.Down()
}

func (g *GarageDoorDownCommand) Undo() {
	g.GarageDoor.Up()
}
