package fly_behavior

import "fmt"

type FlyNoWay struct{}

func (flynoway *FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}
