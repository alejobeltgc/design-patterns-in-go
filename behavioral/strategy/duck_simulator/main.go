package main

import (
	"designpatterns/behavioral/strategy/duck_simulator/fly_behavior"
	"designpatterns/behavioral/strategy/duck_simulator/quack_behavior"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Duck Simulator - Strategy Pattern Demo ===")

	mallard := NewMallardDuck()

	fmt.Println("1. MallardDuck con comportamientos iniciales:")
	mallard.Display()
	mallard.PerformFly()
	mallard.PerformQuack()
	mallard.Swim()

	fmt.Println("\n" + strings.Repeat("-", 50) + "\n")

	fmt.Println("2. RubberDuck con comportamientos apropiados:")
	rubber := NewRubberDuck()
	rubber.Display()
	rubber.PerformFly()
	rubber.PerformQuack()
	rubber.Swim()

	fmt.Println("\n" + strings.Repeat("-", 50) + "\n")

	fmt.Println("3. Cambiando comportamiento en tiempo de ejecución:")
	fmt.Println("Mallard antes del cambio:")
	mallard.Display()
	mallard.PerformFly()
	mallard.PerformQuack()

	mallard.SetFlyBehavior(&fly_behavior.FlyNoWay{})
	mallard.SetQuackBehavior(&quack_behavior.Squeak{})

	fmt.Println("\nMallard después del cambio:")
	mallard.Display()
	mallard.PerformFly()
	mallard.PerformQuack()

	fmt.Println("\n" + strings.Repeat("-", 50) + "\n")

	fmt.Println("4. Creando un pato genérico con comportamientos personalizados:")
	genericDuck := &Duck{}
	genericDuck.SetFlyBehavior(&fly_behavior.FlyWithWings{})
	genericDuck.SetQuackBehavior(&quack_behavior.MuteQuack{})

	genericDuck.Display()
	genericDuck.PerformFly()
	genericDuck.PerformQuack()

	fmt.Println("\n=== Demo completado ===")
}
