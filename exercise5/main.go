package main

import "fmt"

type anupam int

var x anupam
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
