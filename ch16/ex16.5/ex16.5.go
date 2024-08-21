package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5)

	slice2 := append(slice1, 4, 5)
	// cap() 함수를 이용해 슬라이스 capacity 값을 알 수 있습니다.
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100
	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 100)
	slice1 = append(slice1, 200)

	fmt.Println("After append 100, 200")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}
