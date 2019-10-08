package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{3, 5, 2, 99, 23, 23, 1}
	xs := []string{"Anupam", "Sahil", "James", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("-----------------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
