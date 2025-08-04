package concretecommands

import "designpatterns/behavioral/command/remote/devices"

type CeilingFanOffCommand struct {
	ceilingFan *devices.CeilingFan
	prevSpeed  int
}

func NewCeilingFanOffCommand(ceilingFan *devices.CeilingFan) *CeilingFanOffCommand {
	return &CeilingFanOffCommand{
		ceilingFan: ceilingFan,
	}
}

func (c *CeilingFanOffCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.Off()
}

func (c *CeilingFanOffCommand) Undo() {
	switch c.prevSpeed {
	case 0: // OFF
		c.ceilingFan.Off()
	case 1: // LOW
		c.ceilingFan.Low()
	case 2: // MEDIUM
		c.ceilingFan.Medium()
	case 3: // HIGH
		c.ceilingFan.High()
	}
}
