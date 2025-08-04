package commandinterface

type Command interface {
	Execute()
	Undo()
}
