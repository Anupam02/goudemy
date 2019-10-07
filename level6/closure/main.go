package main

import "fmt"

var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

	{
		y := 1
		fmt.Println(y)
	}
	// fmt.Println(y) - won't work y scope is code block above
	fmt.Println("----------------------------")
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
func foo() {
	fmt.Println("Hello")
	x++
}
