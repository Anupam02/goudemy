package main

import "fmt"

func main() {
	fmt.Println("Hello from main")
	s := sum(33, 34)
	fmt.Println(s)
	li := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := sum(li...)
	fmt.Println(s1)
	s2 := even(sum, li...)
	fmt.Println("even numbers sum: ", s2)
	s3 := odd(sum, li...)
	fmt.Println("odd numbers sum: ", s3)

}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 1 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
