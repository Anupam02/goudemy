package main

import "fmt"

func main() {
	var xi [5]int
	for i := range xi {
		xi[i] = i
	}
	for i, v := range xi {
		fmt.Printf("%d\t%d\n", i, v)
	}
	fmt.Printf("%T\n", xi)
}
