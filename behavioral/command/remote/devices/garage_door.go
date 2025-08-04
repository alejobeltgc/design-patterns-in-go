package devices

import "fmt"

type GarageDoor struct{}

func (g *GarageDoor) Up() {
	fmt.Println("Puerta de garage abierta")
}

func (g *GarageDoor) Down() {
	fmt.Println("Puerta de garage cerrada")
}

func (g *GarageDoor) Stop() {
	fmt.Println("Puerta de garage detenida")
}

func (g *GarageDoor) LightOn() {
	fmt.Println("Luz del garage encendida")
}

func (g *GarageDoor) LightOff() {
	fmt.Println("Luz del garage apagada")
}
