package main

import "fmt"

func main() {
	ja := []string{"Joshua", "Aaron", "Chocolate", "Manna"}
	fmt.Println(ja)
	as := []string{"Aaron", "Schultz", "Strawberry", "Bread"}
	fmt.Println(as)

	choir := [][]string{ja, as}
	fmt.Println(choir)
}
