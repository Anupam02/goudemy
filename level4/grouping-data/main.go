package main

import "fmt"

func main() {
	// x := type {values} // composite literals
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(len(x))
	for i, v := range x {
		fmt.Println(i, v)
	}
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
	x = append(x, 77, 88, 99, 1010)
	fmt.Println(x)
	y := []int{34, 234, 2345, 3453}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

// a SLICE allows you to group together VALUES of the same type
