package main

import "fmt"

func main() {
	var x [5]int
	var y [6]int
	fmt.Printf("%v\t%T\n", x, x)
	fmt.Printf("%v\t%T\n", y, y)
	x[3] = 42
	y[3] = 24
	fmt.Printf("%v\t%T\n", x, x)
	fmt.Printf("%v\t%T\n", y, y)
}
