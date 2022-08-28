package command

import (
	"design-pattern/behavioral/command/service"
)

type addCustomerCMD struct {
	customer service.Customer
}

func (c *addCustomerCMD) Execute() {
	c.customer.AddCustomer()
}

func NewAddCustomerCMD(customer service.Customer) *addCustomerCMD {
	return &addCustomerCMD{
		customer: customer,
	}
}
