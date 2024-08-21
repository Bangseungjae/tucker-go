package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}
func (a *account) withdraw(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100} // balance가 100인 account 포인터 변수 생성
	withdrawFunc(a, 30)
	a.withdraw(30)

	fmt.Printf("%d \n", a.balance)
}
