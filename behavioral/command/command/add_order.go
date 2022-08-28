package command

import "design-pattern/behavioral/command/service"

type addOrderCMD struct {
	order service.Order
}

func (c *addOrderCMD) Execute() {
	c.order.AddOrder()
}

func NewAddOrderCMD(order service.Order) *addOrderCMD {
	return &addOrderCMD{
		order: order,
	}
}
