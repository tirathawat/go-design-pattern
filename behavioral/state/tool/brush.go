package tool

import "fmt"

type brush struct{}

func (t *brush) MouseDown() {
	fmt.Println("show brush icon")
}

func (t *brush) MouseUp() {
	fmt.Println("draw a line")
}

func NewBrush() *brush {
	return &brush{}
}
