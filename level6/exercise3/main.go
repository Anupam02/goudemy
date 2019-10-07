package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("greetings from main.")
}

func foo() {
	defer func() {
		fmt.Println("Greetings from Anonymous function in foo.")
	}()
	fmt.Println("greetings from foo.")
}
