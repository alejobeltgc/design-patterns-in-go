package devices

import "fmt"

type PopcornPopper struct {
	Description string
}

func NewPopcornPopper(description string) *PopcornPopper {
	return &PopcornPopper{
		Description: description,
	}
}

func (p *PopcornPopper) On() {
	fmt.Printf("%s está encendida\n", p.Description)
}

func (p *PopcornPopper) Off() {
	fmt.Printf("%s está apagada\n", p.Description)
}

func (p *PopcornPopper) Pop() {
	fmt.Printf("%s haciendo palomitas!\n", p.Description)
}
