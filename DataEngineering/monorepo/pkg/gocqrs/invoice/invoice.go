package invoice

import (
	"api_thienhang_com/pkg/gocqrs/base/aggregate"
	"api_thienhang_com/pkg/gocqrs/base/command"
	"api_thienhang_com/pkg/gocqrs/base/event"
	"api_thienhang_com/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type Invoice struct {
	aggregate.BaseAggregate
	Owner         string // UUID of user
	Method        string // Payment Method
	Amount        int    // The Amount
	Status        int
	Discount      string
	DepositAmount string
}

// ApplyChange to Invoice
func (a *Invoice) ApplyChange(evt event.Event) {
	switch e := evt.Data.(type) {
	case *InvoiceCreated:
		a.Method = e.Method
		a.Amount = e.Amount
		a.Discount = e.Discount
		a.DepositAmount = e.DepositAmount
		a.ID = evt.AggregateID
	case *DepositPerformed:
		a.Amount = e.Amount
		a.Status = e.Status
		a.Discount = e.Discount
		a.DepositAmount = e.DepositAmount
		//a.Owner = e.Owner
	case *WithdrawalPerformed:
		//a.Amount -= e.Amount
	}
}

// HandleCommand create events and validate based on such command
func (a *Invoice) HandleCommand(cmd command.Command) error {
	log.Error(cmd)
	evt := event.Event{
		AggregateID:   a.ID,
		AggregateType: "Invoice",
	}

	switch c := cmd.(type) {
	case CreateInvoice:
		evt.AggregateID = c.AggregateID
		evt.CreatedBy = c.CreatedBy
		evt.CreatedFB = c.CreatedFB
		evt.Data = &InvoiceCreated{
			c.Method,
			c.Amount,
			c.Discount,
			c.DepositAmount,
		}

	case PerformDeposit:
		if a.Amount != c.Amount {
			return utils.ErrAmountOut
		}

		evt.Data = &DepositPerformed{
			c.Status,
			c.Amount,
			c.Discount,
			c.DepositAmount,
			//c.Owner,
		}
	case PerformWithdrawal:
		log.Error("XXXXAAAA")
	}

	a.BaseAggregate.ApplyChangeHelper(a, evt, true)
	return nil
}
