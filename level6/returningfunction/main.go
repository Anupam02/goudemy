package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := func() int {
		return 123
	}()
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y := bar()
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	yVal := y()
	fmt.Println(yVal)
	fmt.Println(bar()())
}

func foo() string {
	s := "Hello world"
	return s
}

func bar() func() int {
	return func() int {
		return 345
	}
}
