package main

import "fmt"

func main() {
	foo(1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("sum := ", sum)
}

/// func ( r receiver ) identifier(parameter(s)) (retun(s)) { code }
