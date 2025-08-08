package devices

import "fmt"

type Amplifier struct {
	Description string
	Volume      int
}

func NewAmplifier(description string) *Amplifier {
	return &Amplifier{
		Description: description,
		Volume:      0,
	}
}

func (a *Amplifier) On() {
	fmt.Printf("%s está encendido\n", a.Description)
}

func (a *Amplifier) Off() {
	fmt.Printf("%s está apagado\n", a.Description)
}

func (a *Amplifier) SetVolume(volume int) {
	a.Volume = volume
	fmt.Printf("%s volumen establecido a %d\n", a.Description, volume)
}

func (a *Amplifier) SetSurroundSound() {
	fmt.Printf("%s configurado para sonido surround\n", a.Description)
}

func (a *Amplifier) SetStereoSound() {
	fmt.Printf("%s configurado para sonido estéreo\n", a.Description)
}
