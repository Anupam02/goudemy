package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hello My name is %v %v , my age is %d.\n", p.first, p.last, p.age)
}

func main() {
	p1 := person{
		first: "Anupam",
		last:  "Patel",
		age:   26,
	}
	p1.speak()
}
