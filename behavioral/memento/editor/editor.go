package editor

type Editor struct {
	Title     string
	Content   string
	FontSize  int
	FontColor string
}

func (e *Editor) Save() state {
	return setState(e)
}

func (e *Editor) Restore(m state) {
	m.getState(e)
}
