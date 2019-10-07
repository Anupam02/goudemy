package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	xs := foo(xi...)
	fmt.Println(xs)
	xt := bar(xi)
	fmt.Println(xt)
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
