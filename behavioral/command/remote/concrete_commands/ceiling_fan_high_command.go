package concretecommands

import "designpatterns/behavioral/command/remote/devices"

type CeilingFanHighCommand struct {
	ceilingFan *devices.CeilingFan
	prevSpeed  int
}

func NewCeilingFanHighCommand(ceilingFan *devices.CeilingFan) *CeilingFanHighCommand {
	return &CeilingFanHighCommand{
		ceilingFan: ceilingFan,
	}
}

func (c *CeilingFanHighCommand) Execute() {
	c.prevSpeed = c.ceilingFan.GetSpeed()
	c.ceilingFan.High()
}

func (c *CeilingFanHighCommand) Undo() {
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
