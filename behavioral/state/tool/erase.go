package tool

import "fmt"

type erase struct{}

func (t *erase) MouseDown() {
	fmt.Println("show eraser icon")
}

func (t *erase) MouseUp() {
	fmt.Println("erase something")
}

func NewErase() *erase {
	return &erase{}
}
