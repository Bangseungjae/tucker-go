package main

import "fmt"

func main() {
	var a = [2][5]int{
		{1, 2, 3, 4, 5},  // b[0] 초기화용 { }
		{6, 7, 8, 9, 10}, // b[1] 초기화용 { }
	}

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
