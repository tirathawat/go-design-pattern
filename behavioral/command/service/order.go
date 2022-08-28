package service

type Order interface {
	AddOrder()
	RemoveOrder()
}

type order struct{}

func (o *order) AddOrder() {
	println("add order")
}

func (o *order) RemoveOrder() {
	println("remove order")
}

func NewOrder() *order {
	return &order{}
}
