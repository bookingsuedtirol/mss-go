package payment_method

import "github.com/HGV/mss-go/shared"

const (
	DepositByCreditCard shared.PaymentMethod = 1 << iota
	CreditCardAsSecurity
	DepositByBankTransfer
	PaymentByCreditCard
	PaymentByBankTransfer
	PaymentAtTheHotel
)
