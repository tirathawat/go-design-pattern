package command

import (
	"design-pattern/behavioral/command/service"
	"fmt"
)

type removeOrderCMD struct {
	order   service.Order
	history history
}

func (c *removeOrderCMD) Execute() {
	c.order.RemoveOrder()
	c.history.Push(c)
}

func (c *removeOrderCMD) UnExecute() {
	fmt.Println("unexecute remove order")
}

func (c *removeOrderCMD) GetHistory() history {
	return c.history
}

func NewRemoveOrderCMD(order service.Order) *removeOrderCMD {
	return &removeOrderCMD{
		order:   order,
		history: history{commands: make([]UndoableCommand, 0), current: 0},
	}
}
