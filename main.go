package main

import (
	"interfaces/payments"
	"interfaces/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewCrypto()

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("Burger", 5)
	paymentModule.Pay("Pizza", 10)
	paymentModule.Pay("Coke", 3)

	allInfo := paymentModule.AllInfo()

	pp.Println("Все оплаты:", allInfo)
	// Здесь можно инициализировать и использовать PaymentModule с разными методами оплаты.
}
