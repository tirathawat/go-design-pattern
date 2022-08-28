package main

import "design-pattern/behavioral/template_method/transaction"

func main() {
	t := transaction.New()
	t.TransferMoney()
}
