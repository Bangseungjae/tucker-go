package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

// 포인터 메서드
func (a *account) withdrawPointer(amount int) {
	a.balance -= amount
}

// 값 타입 메서드
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// 변경된 값을 반환하는 값 타입 메서드
func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{100, "Joe", "Park"} // balance가 100인 account 포인터 변수 생성
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance) // 70

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance) // 70

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance) // 50

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance) // 20
}
