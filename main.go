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
	idPhone := paymentModule.Pay("Phone", 500)
	idCoke := paymentModule.Pay("Coke", 3)

	paymentModule.Cancel(idPhone)

	allInfo := paymentModule.AllInfo()

	pp.Println("Все оплаты:", allInfo)

	cokeInfo := paymentModule.Info(idCoke)
	pp.Println("Coke Info:", cokeInfo)
	// Здесь можно инициализировать и использовать PaymentModule с разными методами оплаты.
}
