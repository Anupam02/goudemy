package main

import "fmt"

func main() {
	s := struct {
		first string
		last  string
		age   int
	}{
		first: "Anupam",
		last:  "Patel",
		age:   27,
	}
	fmt.Println(s)
}
