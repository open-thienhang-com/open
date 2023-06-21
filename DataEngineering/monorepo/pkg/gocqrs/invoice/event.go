package invoice

// InvoiceCreated event
type InvoiceCreated struct {
	Method        string `json:"method"`
	Amount        int    `json:"amount"`
	Discount      string `json:"discount"`
	DepositAmount string `json:"deposit_amount"`
}

// DepositPerformed event
type DepositPerformed struct {
	Status        int    `json:"status"`
	Amount        int    `json:"amount"`
	Discount      string `json:"discount"`
	DepositAmount string `json:"deposit_amount"`
}

// WithdrawalPerformed event
type WithdrawalPerformed struct {
	Amount int `json:"ammount"`
}
