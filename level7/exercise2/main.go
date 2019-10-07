package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "Anupam"
	p.last = "Patel"
	p.age = 26
}

func main() {
	p := person{
		first: "Sahil",
		last:  "Virmani",
		age:   26,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
