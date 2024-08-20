package main

import "fmt"

func main() {
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var slice1 []int = array[1:3:5]
	fmt.Println("slice1:", slice1, "cap: ", cap(slice1))
}
