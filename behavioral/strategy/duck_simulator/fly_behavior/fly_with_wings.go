package fly_behavior

import "fmt"

type FlyWithWings struct{}

func (flywithwings *FlyWithWings) Fly() {
	fmt.Println("I'm flying with wings!")
}
