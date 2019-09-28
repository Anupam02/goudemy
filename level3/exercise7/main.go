package main

import "fmt"

func main() {
	x := -1
	if x > 4 {
		fmt.Println("x is greater than 4")
	} else if x > 0 {
		fmt.Println("x is greater than 0")
	} else {
		fmt.Println("x is a non positive number")
	}
}
