package command

type composite struct {
	cmds []Command
}

func (c *composite) Execute() {
	for _, cmd := range c.cmds {
		cmd.Execute()
	}
}

func NewComposite(cmds []Command) *composite {
	return &composite{
		cmds: cmds,
	}
}
