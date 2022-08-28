package transaction

import (
	"design-pattern/behavioral/template_method/audit"
	"fmt"
)

type transaction struct {
	audit audit.Audit
}

func (t *transaction) TransferMoney() {
	t.audit.Execute(func() {
		fmt.Println("transfer money")
	})
}

func New() *transaction {
	return &transaction{
		audit: audit.New(),
	}
}
