package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am a secretAgent ", s.first, s.last)
}
func (p person) speak() {
	fmt.Println("I am a person ", p.first, p.last)
}

// interfaces allows us to define behaviour and do polymorphism
// keyword identifier type(int, struct, interface etc..)
// anyone who speack() is also a human
// a value can be of more than one type
type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first)
	default:
		fmt.Println("default")
	}
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()
	fmt.Println(p1)
	p1.speak()
	bar(sa1)
	bar(sa2)
	bar(p1)
	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(y)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
