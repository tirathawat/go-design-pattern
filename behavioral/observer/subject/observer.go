package subject

import "fmt"

type Listeners map[string]*Listener

type observer interface {
	onUpdate(data string)
}

type Listener struct {
	Name string
}

func (l *Listener) onUpdate(data string) {
	fmt.Println("Listener:", l.Name, "got data change", data)
}

func (ls Listeners) add(l *Listener) {
	ls[l.Name] = l
}

func (ls Listeners) remove(l *Listener) {
	delete(ls, l.Name)
}
