package devices

import "fmt"

const (
	OFF = iota
	LOW
	MEDIUM
	HIGH
)

type CeilingFan struct {
	speed int
}

func NewCeilingFan() *CeilingFan {
	return &CeilingFan{speed: OFF}
}

func (c *CeilingFan) High() {
	c.speed = HIGH
	fmt.Println("Ventilador de techo está en velocidad ALTA")
}

func (c *CeilingFan) Medium() {
	c.speed = MEDIUM
	fmt.Println("Ventilador de techo está en velocidad MEDIA")
}

func (c *CeilingFan) Low() {
	c.speed = LOW
	fmt.Println("Ventilador de techo está en velocidad BAJA")
}

func (c *CeilingFan) Off() {
	c.speed = OFF
	fmt.Println("Ventilador de techo está APAGADO")
}

func (c *CeilingFan) GetSpeed() int {
	return c.speed
}
