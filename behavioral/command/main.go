package main

import (
	"design-pattern/behavioral/command/button"
	"design-pattern/behavioral/command/command"
	"design-pattern/behavioral/command/service"
)

func main() {
	b := button.New("add customer", command.NewAddCustomerCMD(service.NewCustomer()))
	b.Click()

	b = button.New("add order", command.NewAddOrderCMD(service.NewOrder()))
	b.Click()

	b = button.New("add customer and order", command.NewComposite([]command.Command{
		command.NewAddCustomerCMD(service.NewCustomer()),
		command.NewAddOrderCMD(service.NewOrder()),
	}))
	b.Click()

	removeOrderCommand := command.NewRemoveOrderCMD(service.NewOrder())
	b = button.New("remove order", removeOrderCommand)
	b.Click()

	b = button.New("undo", command.NewUndo(removeOrderCommand.GetHistory()))
	b.Click()
}
