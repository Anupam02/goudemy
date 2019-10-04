package main

import "fmt"

type person struct {
	first                     string
	last                      string
	favouriteIceCreamFlavours []string
}

func main() {
	p1 := person{
		first:                     "Anupam",
		last:                      "Patel",
		favouriteIceCreamFlavours: []string{"Vanilla", "Butterscotch"},
	}
	p2 := person{
		first:                     "Sahil",
		last:                      "Virmani",
		favouriteIceCreamFlavours: []string{"Strawberry", "Mango"},
	}
	fmt.Printf("%v %v likes : ", p1.first, p1.last)
	for _, v := range p1.favouriteIceCreamFlavours {
		fmt.Printf("%v ", v)
	}
	fmt.Println("")
	fmt.Printf("%v %v likes : ", p2.first, p2.last)
	for _, v := range p2.favouriteIceCreamFlavours {
		fmt.Printf("%v ", v)
	}
}
