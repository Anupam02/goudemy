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
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for _, v := range m {
		fmt.Printf("%v %v's favourite ice cream flavours are : ", v.first, v.last)
		for _, fav := range v.favouriteIceCreamFlavours {
			fmt.Printf("%v ", fav)
		}
		fmt.Println("")
	}
}
