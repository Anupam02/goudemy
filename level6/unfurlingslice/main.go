package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum("James", xi...)
	y := sum("Anupam")
	fmt.Println("The total is ", x)
	fmt.Println("The total is ", y)
}

// ... (variadic) means zero or more
func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
