package main

import (
	"design-pattern/creation/builder/notification"
	"fmt"
)

func main() {
	notification, err := notification.NewBuilder().
		SetTitle("New Notification").
		SetIcon("icon.png").
		SetImage("image.png").
		SetPriority(10).
		SetMessage("This is a new notification").
		Build()

	if err != nil {
		panic(err)
	}

	fmt.Println(notification.String())
}
