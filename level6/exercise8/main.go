package main

import "fmt"

func main() {
	r := foo()
	fmt.Println(r())

}

func foo() func() int {
	return func() int {
		return 42
	}
}
