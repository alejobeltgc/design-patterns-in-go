package concretecommands

import commandinterface "designpatterns/behavioral/command/remote/command_interface"

type MacroCommand struct {
	commands []commandinterface.Command
}

func NewMacroCommand(commands []commandinterface.Command) *MacroCommand {
	return &MacroCommand{
		commands: commands,
	}
}

func (m *MacroCommand) Execute() {
	for _, command := range m.commands {
		command.Execute()
	}
}

func (m *MacroCommand) Undo() {
	for i := len(m.commands) - 1; i >= 0; i-- {
		m.commands[i].Undo()
	}
}
