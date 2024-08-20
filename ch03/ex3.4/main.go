package main

import "fmt"

func main() {
	str := "Hello\tGo\t\tWorld\"Go\" is Awesome"

	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str)

}
