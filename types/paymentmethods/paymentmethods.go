package paymentmethods

import "github.com/HGV/mss-go/response"

const (
	CreditCard response.PaymentMethods = 8 << iota
	ATM        response.PaymentMethods = 32 << iota
	Mastercard
	Visa
	DinersClub
	AmericanExpress
)
