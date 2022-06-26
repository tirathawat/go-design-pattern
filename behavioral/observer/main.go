package main

import "design-pattern/behavioral/observer/subject"

func main() {
	listener1 := subject.Listener{
		Name: "_listener1",
	}

	listener2 := subject.Listener{
		Name: "_listener2",
	}

	subject := subject.New()
	subject.Register(&listener1)
	subject.Register(&listener2)

	subject.ChangeItem("Monday")

	subject.Unregister(&listener1)

	subject.ChangeItem("Sunday")
}
