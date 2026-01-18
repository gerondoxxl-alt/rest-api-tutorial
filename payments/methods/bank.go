package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (c Bank) Pay(usd int) int {
	fmt.Println("Оплата банковским переводом!")
	fmt.Println("Размер оплаты:", usd, "долларов")

	return rand.Int()

}

func (c Bank) Cancel(id int) {
	fmt.Println("Отмена оплаты банковским переводом!")
	fmt.Println("ID отменяемой операции:", id)
}
