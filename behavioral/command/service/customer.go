package service

import "fmt"

type Customer interface {
	AddCustomer()
}

type customer struct{}

func (c *customer) AddCustomer() {
	fmt.Println("Add Customer")
}

func NewCustomer() *customer {
	return &customer{}
}
