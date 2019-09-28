package main

import "t"

func main() {
	switch {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("True")
	default:
		fmt.Println("default")
	}
}
