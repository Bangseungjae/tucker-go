package main

import "fmt"

// 포인터를 사용하지 않는 예

type Data struct {
	Value int
	Data  [200]int
}

func ChangeData(arg Data) {
	arg.Value = 999
	arg.Data[100] = 999
}

func main() {
	var data Data

	ChangeData(data)
	fmt.Printf("Value = %d\n", data.Value)
	fmt.Printf("data[100] = %d\n", data.Data[100])
}
