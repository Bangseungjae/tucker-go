package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera"}
	fmt.Println(names)
	n, found := slices.BinarySearch(names, "Vera")
	fmt.Println("Vera:", n, found)
	n, found = slices.BinarySearch(names, "Bill")
	fmt.Println("Bill:", n, found)
	if !found {
		fmt.Println("Nothing found")
		names = slices.Insert(names, n, "Bill")
	}
	fmt.Println(names)
}
