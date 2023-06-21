package invoice

import (
	"api_thienhang_com/pkg/gocqrs/base/command"
)

// CreateInvoice assigned to an owner
type CreateInvoice struct {
	command.BaseCommand
	//Owner     string
	Method        string
	Amount        int
	CreatedBy     string
	CreatedFB     string
	Discount      string `json:"discount"`
	DepositAmount string `json:"deposit_amount"`
}

// PerformDeposit to a given Invoice
type PerformDeposit struct {
	command.BaseCommand
	Status        int
	Amount        int
	Discount      string `json:"discount"`
	DepositAmount string `json:"deposit_amount"`
}

// PerformWithdrawal to a given Invoice
type PerformWithdrawal struct {
	command.BaseCommand
	Amount int
}
