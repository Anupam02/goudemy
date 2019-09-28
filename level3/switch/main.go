package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case !true:
		fmt.Println("this is true")
	case (3 == 5):
		fmt.Println("3 != 5")
	case false:
		fmt.Println("this is false case after fall through")
	default:
		fmt.Println("this is default")
	}
}
