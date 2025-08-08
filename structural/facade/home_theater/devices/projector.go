package devices

import "fmt"

type Projector struct {
	Description string
	Input       string
}

func NewProjector(description string) *Projector {
	return &Projector{
		Description: description,
		Input:       "",
	}
}

func (p *Projector) On() {
	fmt.Printf("%s está encendido\n", p.Description)
}

func (p *Projector) Off() {
	fmt.Printf("%s está apagado\n", p.Description)
}

func (p *Projector) SetInput(input string) {
	p.Input = input
	fmt.Printf("%s entrada establecida a %s\n", p.Description, input)
}

func (p *Projector) WideScreenMode() {
	fmt.Printf("%s en modo pantalla ancha (16:9)\n", p.Description)
}

func (p *Projector) TVMode() {
	fmt.Printf("%s en modo TV (4:3)\n", p.Description)
}
