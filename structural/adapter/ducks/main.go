package main

import (
	"designpatterns/structural/adapter/ducks/adaptee"
	"designpatterns/structural/adapter/ducks/adapter"
	"designpatterns/structural/adapter/ducks/interfaces"
	"fmt"
)

func main() {
	fmt.Println("=== Adapter Pattern Demo - Patos y Pavos ===")

	// Crear instancias de los objetos originales
	duck := &adaptee.MallardDuck{}
	turkey := &adaptee.WildTurkey{}

	fmt.Println("\n1. Comportamiento original del pato:")
	testDuck(duck)

	fmt.Println("\n2. Comportamiento original del pavo:")
	testTurkey(turkey)

	fmt.Println("\n3. Problema: ¿Qué pasa si queremos usar un pavo donde esperamos un pato?")
	fmt.Println("   No podemos hacer: testDuck(turkey) - ¡Error de compilación!")

	fmt.Println("\n4. Solución: Usar TurkeyAdapter para adaptar Turkey a Duck")
	turkeyAdapter := adapter.NewTurkeyAdapter(turkey)
	fmt.Println("   Ahora podemos usar el pavo como si fuera un pato:")
	testDuck(turkeyAdapter)

	fmt.Println("\n5. Demostración bidireccional: DuckAdapter para usar Duck como Turkey")
	duckAdapter := adapter.NewDuckAdapter(duck)
	fmt.Println("   Ahora podemos usar el pato como si fuera un pavo:")
	testTurkey(duckAdapter)

	fmt.Println("\n=== Comparación lado a lado ===")
	fmt.Println("\nPato original:")
	testDuck(duck)

	fmt.Println("\nPavo adaptado como pato:")
	testDuck(turkeyAdapter)

	fmt.Println("\nPavo original:")
	testTurkey(turkey)

	fmt.Println("\nPato adaptado como pavo:")
	testTurkey(duckAdapter)

	fmt.Println("\n=== Demo completado ===")
}

func testDuck(duck interfaces.Duck) {
	fmt.Print("  ")
	duck.Quack()
	fmt.Print("  ")
	duck.Fly()
}

func testTurkey(turkey interfaces.Turkey) {
	fmt.Print("  ")
	turkey.Gobble()
	fmt.Print("  ")
	turkey.Fly()
}
