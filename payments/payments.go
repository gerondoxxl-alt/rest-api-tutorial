package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) PaymentModule {
	return PaymentModule{
		paymentMethod: paymentMethod,
	}
}

//Метод Pay()
//Принимает:
// 1. Описание проводимой оплаты.
// 2. Сумма оплаты.
// Возвращает:
// 1. ID проведенной операции.

func (p PaymentModule) Pay(description string, usd int) int {
	id := p.paymentMethod.Pay(usd)

	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}
	// 1. Проводить оплату
	// 2. Получать ID проведенной оплаты.
	// 3. Сохранять информацию о проведенной операции.
	// - описание операции
	// - сколько было потрачено
	// - отмененная ли операция
	// 4. Возвращать ID проведенной операции.

}

// Метод Cancel()
// Принимает:
// 1. ID проведеной операции.
// Возвращает:
// Ничего не возвращает.
func (p PaymentModule) Cancel() {}

// Метод Info()
// Принимает:
// 1. ID операции.
// Возвращает:
// Иноформацию о проведенной операции
func (p PaymentModule) Info() {}

// Метод AllInfo()
// Принимает:
// - ничего не принимает
// Возвращает:
// 1. Информацию о всех проведенных операциях
func (p PaymentModule) AllInfo() {}
