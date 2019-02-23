//Но самым идеалогически правильным путем будет обеспечение синхронизации, используя каналы!
// Don't communicate by sharing memory; share memory by communicating

package main

import (
	"fmt"
	"github.com/pkg/errors"
)

// Аккаунт - это структура, содержащая в себе каналы для изменений
type AccountAsync struct {
	balance     float64
	deltaChan   chan float64
	balanceChan chan float64
	errorChan   chan error
}

//Потребуется конструктор, чтобы скрыть от пользователя тонкости реализации
func NewAccount(balance float64) (a *AccountAsync) {
	a = &AccountAsync{
		balance:     balance,
		deltaChan:   make(chan float64),
		balanceChan: make(chan float64),
		errorChan:   make(chan error),
	}
	//и запустим горутину для обслуживания операций с параметром
	go a.run()
	return
}

func (a *AccountAsync) Balance() float64 {
	return <-a.balanceChan
}

//записываем количество в канал изменений
func (a *AccountAsync) Deposit(amount float64) error {
	a.deltaChan <- amount
	return <-a.errorChan
}

//Аналогично , по сути, это функция нужна только для сохранения семантики
func (a *AccountAsync) Withdraw(amount float64) error {
	a.deltaChan <- amount
	return <-a.errorChan
}

//Применение изменений к счет
func (a *AccountAsync) applyDelta(amount float64) error {
	stateStr := "Кладем в счет"
	if amount < 0 {
		stateStr = "Снимаем"
	}
	fmt.Println(stateStr, amount)

	newBalance := a.balance + amount
	if newBalance < 0 {
		return errors.New("Insufficient funds")
	}
	a.balance = newBalance
	return nil
}

/*
Бесконечный цикл обработки счетчика
Теперь сколько бы горутин не производили операции над этим аккаунтом,
все они будут синхронизированы здесь, и блокировки уже не нужны
*/
func (a *AccountAsync) run() {
	var delta float64
	for {
		select {
		//Если поступили изменения
		case delta = <-a.deltaChan:
			//Попробуем их применить
			a.errorChan <- a.applyDelta(delta)
		//Если кто-то запрашивает баланс
		case a.balanceChan <- a.balance:
			//не делаем ничего так как мы уже отправили ответ
		}
	}
}

func main() {
	acc := NewAccount(20)

	//Стартуем 10 горутин
	for i := 0; i < 10; i++ {
		go func() {
			//Каждая из которых производит операции с аккаунтом
			for j := 0; j < 10; j++ {
				//Иногда снимает деньги
				if j%2 == i {
					acc.Withdraw(50)
					continue
				}
				//иногда кладет
				acc.Deposit(50)
			}
		}()
	}
	fmt.Scanln()
	//Теперь баланс всегда будет сходиться в ноль
	fmt.Println(acc.balance)
}
