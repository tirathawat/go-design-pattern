package audit

import "fmt"

type Audit interface {
	record()
	stop()
	Execute(fn func())
}

type audit struct{}

func (a *audit) record() {
	fmt.Println("start logging")
}

func (a *audit) stop() {
	fmt.Println("stop logging")
}

func (a *audit) Execute(fn func()) {
	a.record()
	fn()
	a.stop()
}

func New() *audit {
	return &audit{}
}
