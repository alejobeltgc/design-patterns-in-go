package main

import (
	"designpatterns/structural/facade/home_theater/devices"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Facade Pattern Demo - Home Theater ===")

	// Crear todos los componentes del subsistema complejo
	amp := devices.NewAmplifier("Amplificador Top-O-Line")
	tuner := devices.NewTuner("Sintonizador AM/FM Top-O-Line")
	dvd := devices.NewDVDPlayer("Reproductor DVD Top-O-Line")
	projector := devices.NewProjector("Proyector Top-O-Line")
	screen := devices.NewScreen("Pantalla de Teatro")
	lights := devices.NewTheaterLights("Luces del Teatro")
	popper := devices.NewPopcornPopper("Máquina de Palomitas")

	// Crear la fachada que simplifica el uso del sistema complejo
	homeTheater := NewHomeTheaterFacade(amp, tuner, dvd, projector, screen, lights, popper)

	fmt.Println("\n=== Sin Facade: Configuración manual compleja ===")
	fmt.Println("Para ver una película manualmente necesitarías:")
	fmt.Println("1. Encender máquina de palomitas y hacer palomitas")
	fmt.Println("2. Atenuar las luces al 10%")
	fmt.Println("3. Bajar la pantalla")
	fmt.Println("4. Encender proyector y configurar modo pantalla ancha")
	fmt.Println("5. Encender amplificador, configurar surround y volumen")
	fmt.Println("6. Encender DVD y reproducir película")
	fmt.Println("¡Son muchos pasos!")

	fmt.Println("\n=== Con Facade: Interfaz simplificada ===")
	fmt.Println("Con la fachada, solo necesitas un método:")

	// Usar la fachada para ver una película (simplifica todo el proceso)
	homeTheater.WatchMovie("El Padrino")

	fmt.Println("\n" + strings.Repeat("=", 50))

	// Terminar la película
	homeTheater.EndMovie()

	fmt.Println("\n=== Otra funcionalidad simplificada ===")

	// Usar la fachada para escuchar radio
	homeTheater.ListenToRadio(101.5)

	fmt.Println("\n" + strings.Repeat("=", 30))

	// Terminar la radio
	homeTheater.EndRadio()

	fmt.Println("\n=== Comparación de complejidad ===")
	fmt.Println("Sin Facade:")
	fmt.Println("  - Cliente debe conocer todos los subsistemas")
	fmt.Println("  - Muchas llamadas de métodos")
	fmt.Println("  - Orden específico de operaciones")
	fmt.Println("  - Código duplicado para tareas comunes")

	fmt.Println("\nCon Facade:")
	fmt.Println("  - Una sola llamada de método")
	fmt.Println("  - Interfaz simple y clara")
	fmt.Println("  - Encapsula la complejidad")
	fmt.Println("  - Reutilizable y mantenible")

	fmt.Println("\n=== Demo completado ===")
}
