package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Anupam",
		last:  "Patel",
		age:   26,
	}
	fmt.Println(p1)
}
