package main

import "fmt"

func main() {
	a := 10
	if a == 10 {
		fmt.Println("tested ==")
	}
	if a <= 10 {
		fmt.Println("tested <=")
	}
	if a >= 10 {
		fmt.Println("tested >=")
	}
	if a != 1 {
		fmt.Println("tested !=")
	}
	if a < 11 {
		fmt.Println("tested <")
	}
	if a > 9 {
		fmt.Println("tested >")
	}
}
