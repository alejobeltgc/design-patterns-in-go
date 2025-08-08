package devices

import "fmt"

type Tuner struct {
	Description string
	Frequency   float64
}

func NewTuner(description string) *Tuner {
	return &Tuner{
		Description: description,
		Frequency:   0.0,
	}
}

func (t *Tuner) On() {
	fmt.Printf("%s está encendido\n", t.Description)
}

func (t *Tuner) Off() {
	fmt.Printf("%s está apagado\n", t.Description)
}

func (t *Tuner) SetFrequency(frequency float64) {
	t.Frequency = frequency
	fmt.Printf("%s sintonizado a %.1f FM\n", t.Description, frequency)
}

func (t *Tuner) SetAM() {
	fmt.Printf("%s configurado para AM\n", t.Description)
}

func (t *Tuner) SetFM() {
	fmt.Printf("%s configurado para FM\n", t.Description)
}
