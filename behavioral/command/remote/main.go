package main

import (
	commandinterface "designpatterns/behavioral/command/remote/command_interface"
	concretecommands "designpatterns/behavioral/command/remote/concrete_commands"
	"designpatterns/behavioral/command/remote/devices"
	"designpatterns/behavioral/command/remote/invoker"
	"fmt"
)

func main() {
	fmt.Println("=== Command Pattern Demo - Control Remoto ===")

	// Crear dispositivos (receivers)
	livingRoomLight := devices.NewLight("sala")
	kitchenLight := devices.NewLight("cocina")
	garageDoor := &devices.GarageDoor{}
	ceilingFan := devices.NewCeilingFan()

	// Crear comandos básicos
	lightOn := concretecommands.NewLightOnCommand(livingRoomLight)
	lightOff := concretecommands.NewLightOffCommand(livingRoomLight)
	kitchenLightOn := concretecommands.NewLightOnCommand(kitchenLight)
	kitchenLightOff := concretecommands.NewLightOffCommand(kitchenLight)
	garageUp := concretecommands.NewGarageDoorOpenCommand(garageDoor)
	garageDown := concretecommands.NewGarageDoorDownCommand(garageDoor)
	fanHigh := concretecommands.NewCeilingFanHighCommand(ceilingFan)
	fanOff := concretecommands.NewCeilingFanOffCommand(ceilingFan)

	// Crear macro comando "Party Mode"
	partyOnMacro := concretecommands.NewMacroCommand([]commandinterface.Command{
		lightOn,
		kitchenLightOn,
		fanHigh,
	})

	partyOffMacro := concretecommands.NewMacroCommand([]commandinterface.Command{
		lightOff,
		kitchenLightOff,
		fanOff,
	})

	// Crear control remoto (invoker)
	remote := invoker.NewRemoteControl()

	// Mostrar estado inicial
	fmt.Println("Estado inicial del control remoto:")
	fmt.Println(remote.String())

	// Configurar comandos en diferentes slots
	remote.SetCommand(0, lightOn, lightOff)               // Luz sala
	remote.SetCommand(1, kitchenLightOn, kitchenLightOff) // Luz cocina
	remote.SetCommand(2, fanHigh, fanOff)                 // Ventilador
	remote.SetCommand(3, garageUp, garageDown)            // Garage
	remote.SetCommand(6, partyOnMacro, partyOffMacro)     // Party Mode

	// Mostrar configuración
	fmt.Println("Después de configurar comandos:")
	fmt.Println(remote.String())

	// Probar comandos básicos
	fmt.Println("=== Probando comandos básicos ===")
	fmt.Println("1. Encender luz de sala:")
	remote.OnButtonWasPressed(0)

	fmt.Println("\n2. Encender luz de cocina:")
	remote.OnButtonWasPressed(1)

	fmt.Println("\n3. Deshacer último comando (apagar luz cocina):")
	remote.UndoButtonWasPressed()

	fmt.Println("\n=== Probando comando con estado complejo (ventilador) ===")
	fmt.Println("4. Ventilador a velocidad alta:")
	remote.OnButtonWasPressed(2)

	fmt.Println("\n5. Apagar ventilador:")
	remote.OffButtonWasPressed(2)

	fmt.Println("\n6. Deshacer (volver a velocidad alta):")
	remote.UndoButtonWasPressed()

	fmt.Println("\n7. Deshacer otra vez (volver a OFF):")
	remote.UndoButtonWasPressed()

	fmt.Println("\n=== Probando Macro Command ===")
	fmt.Println("8. Activar 'Party Mode' (macro - enciende todo):")
	remote.OnButtonWasPressed(6)

	fmt.Println("\n9. Deshacer 'Party Mode' (macro undo - apaga todo en orden inverso):")
	remote.UndoButtonWasPressed()

	fmt.Println("\n10. Activar 'Party Mode' otra vez:")
	remote.OnButtonWasPressed(6)

	fmt.Println("\n11. Desactivar 'Party Mode' (macro off):")
	remote.OffButtonWasPressed(6)

	fmt.Println("\n=== Probando comandos de garage ===")
	fmt.Println("12. Abrir garage:")
	remote.OnButtonWasPressed(3)

	fmt.Println("\n13. Deshacer (cerrar garage):")
	remote.UndoButtonWasPressed()

	fmt.Println("\n=== Probando casos especiales ===")
	// Probar slot vacío (NoCommand)
	fmt.Println("14. Presionar botón ON slot 5 (vacío):")
	remote.OnButtonWasPressed(5)

	fmt.Println("\n15. Deshacer después de NoCommand:")
	remote.UndoButtonWasPressed()

	// Probar índice inválido
	fmt.Println("\n16. Presionar botón ON slot 10 (inválido):")
	remote.OnButtonWasPressed(10)

	fmt.Println("\n=== Estado final del control remoto ===")
	fmt.Println(remote.String())

	fmt.Println("\n=== Demo completado ===")
}
