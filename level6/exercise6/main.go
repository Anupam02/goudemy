package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is an Anonymous Function.")
	}()
	fmt.Println("greetings from main")
}
