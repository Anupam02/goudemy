package main

import "fmt"

func main() {
	fmt.Println("hello from main")
	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	f1 := func(x int) {
		fmt.Println("the year big brother started watching:", x)
	}
	f1(1984)
}
