package main

import "fmt"

type anupam int

func main() {
	var x anupam
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
