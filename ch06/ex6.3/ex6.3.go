package main

import "fmt"

const (
	BitFlag1 uint = 1 << iota
	BitFlag2
	BitFlag3
	BitFlag4
)

func main() {
	const (
		C1 = iota + 1
		C2
		C3
	)

	fmt.Println(C1, C2, C3)
	fmt.Println(BitFlag1, BitFlag2, BitFlag3, BitFlag4)

}
