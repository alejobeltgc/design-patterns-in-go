package devices

import "fmt"

type Screen struct {
	Description string
	Position    string
}

func NewScreen(description string) *Screen {
	return &Screen{
		Description: description,
		Position:    "up",
	}
}

func (s *Screen) Up() {
	s.Position = "up"
	fmt.Printf("%s subiendo\n", s.Description)
}

func (s *Screen) Down() {
	s.Position = "down"
	fmt.Printf("%s bajando\n", s.Description)
}
