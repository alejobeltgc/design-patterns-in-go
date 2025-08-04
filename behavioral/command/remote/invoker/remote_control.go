package invoker

import (
	commandinterface "designpatterns/behavioral/command/remote/command_interface"
	concretecommands "designpatterns/behavioral/command/remote/concrete_commands"
	"fmt"
)

type RemoteControl struct {
	OnCommand   [7]commandinterface.Command
	OffCommand  [7]commandinterface.Command
	UndoCommand commandinterface.Command
}

func NewRemoteControl() *RemoteControl {
	noCommand := concretecommands.NewNoCommand()
	rc := &RemoteControl{}

	for i := range 7 {
		rc.OnCommand[i] = noCommand
		rc.OffCommand[i] = noCommand
	}
	rc.UndoCommand = noCommand

	return rc
}

func (s *RemoteControl) SetCommand(slot int, onCommand commandinterface.Command, offCommand commandinterface.Command) {
	if slot < 0 || slot >= 7 {
		fmt.Printf("Error: Slot %d inválido. Debe estar entre 0 y 6\n", slot)
		return
	}

	s.OnCommand[slot] = onCommand
	s.OffCommand[slot] = offCommand
}

func (s *RemoteControl) OnButtonWasPressed(slot int) {
	if slot < 0 || slot >= 7 {
		fmt.Printf("Error: Slot %d inválido. Debe estar entre 0 y 6\n", slot)
		return
	}

	s.OnCommand[slot].Execute()
	s.UndoCommand = s.OnCommand[slot]
}

func (s *RemoteControl) OffButtonWasPressed(slot int) {
	if slot < 0 || slot >= 7 {
		fmt.Printf("Error: Slot %d inválido. Debe estar entre 0 y 6\n", slot)
		return
	}

	s.OffCommand[slot].Execute()
	s.UndoCommand = s.OffCommand[slot]
}

func (s *RemoteControl) UndoButtonWasPressed() {
	s.UndoCommand.Undo()
}

func (s *RemoteControl) String() string {
	result := "\n------ Remote Control -------\n"
	slotNames := []string{
		"Luz Sala    ",
		"Luz Cocina  ",
		"Ventilador  ",
		"Garage      ",
		"Slot 4      ",
		"Slot 5      ",
		"Party Mode  ",
	}

	for i := range 7 {
		onCommand := s.getCommandName(s.OnCommand[i])
		offCommand := s.getCommandName(s.OffCommand[i])
		result += fmt.Sprintf("[%s] %-15s %-15s\n", slotNames[i], onCommand, offCommand)
	}

	undoCommand := s.getCommandName(s.UndoCommand)
	result += fmt.Sprintf("[Undo      ] %s\n", undoCommand)
	result += "-----------------------------\n"

	return result
}

func (s *RemoteControl) getCommandName(cmd commandinterface.Command) string {
	switch cmd.(type) {
	case *concretecommands.LightOnCommand:
		return "Encender"
	case *concretecommands.LightOffCommand:
		return "Apagar"
	case *concretecommands.CeilingFanHighCommand:
		return "Fan Alto"
	case *concretecommands.CeilingFanOffCommand:
		return "Fan Off"
	case *concretecommands.GarageDoorOpenCommand:
		return "Abrir"
	case *concretecommands.GarageDoorDownCommand:
		return "Cerrar"
	case *concretecommands.MacroCommand:
		return "Macro"
	case *concretecommands.NoCommand:
		return "---"
	default:
		return "Desconocido"
	}
}
