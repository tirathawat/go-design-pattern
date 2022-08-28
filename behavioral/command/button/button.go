package button

import "design-pattern/behavioral/command/command"

type button struct {
	text    string
	onclick command.Command
}

func (b *button) Click() {
	b.onclick.Execute()
}

func New(text string, onclick command.Command) *button {
	return &button{
		text:    text,
		onclick: onclick,
	}
}
