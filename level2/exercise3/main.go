package main

import "fmt"

const (
	a = 25
	b = 25.34532
	c = "Anupam Patel"
)

const (
	x int     = 25
	y float64 = 25.9583948
	z string  = "Sahil Virmani"
)

func main() {
	fmt.Printf("a = %v with type %T\n", a, a)
	fmt.Printf("b = %v with type %T\n", b, b)
	fmt.Printf("c = %v with type %T\n", c, c)
	fmt.Printf("x = %v with type %T\n", x, x)
	fmt.Printf("y = %v with type %T\n", y, y)
	fmt.Printf("z = %v with type %T\n", z, z)
}
