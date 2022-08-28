package command

type undoCommand struct {
	history history
}

func (c *undoCommand) Execute() {
	lastCommand := c.history.Pop()
	if lastCommand != nil {
		lastCommand.UnExecute()
	}
}

func NewUndo(history history) *undoCommand {
	return &undoCommand{
		history: history,
	}
}
