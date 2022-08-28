package command

type history struct {
	commands []UndoableCommand
	current  int
}

func (h *history) Push(cmd UndoableCommand) {
	h.commands = append(h.commands, cmd)
	h.current++
}

func (h *history) Pop() UndoableCommand {
	if h.current == 0 {
		return nil
	}
	h.current--
	return h.commands[h.current]
}
