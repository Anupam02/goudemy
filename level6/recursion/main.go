package main

import "fmt"

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	n := factorial(5)
	fmt.Println(n)
	n1 := fact(4)
	fmt.Println(n1)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fact(n int) int {
	result := 1
	for i := 1; i <= n; {
		result *= i
		i++
	}
	return result
}
