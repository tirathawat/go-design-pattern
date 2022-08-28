package tool

import "fmt"

type selection struct {
}

func (t *selection) MouseDown() {
	fmt.Println("show selection icon")
}

func (t *selection) MouseUp() {
	fmt.Println("draw a rectangle")
}

func NewSelection() *selection {
	return &selection{}
}
