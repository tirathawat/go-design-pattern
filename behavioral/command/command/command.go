package command

type Command interface {
	Execute()
}

type UndoableCommand interface {
	Command
	UnExecute()
	GetHistory() history
}
