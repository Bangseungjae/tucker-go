package main

import "fmt"

// 포인터를 사용하지 예

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer)
	fmt.Println(*userPointer)
}
