package main

import "fmt"

type circle struct {
	radius float64
}
type square struct {
	length  float64
	breadth float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (s *square) area() float64 {
	return s.length * s.breadth
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{
		radius: 3,
	}
	info(c)
	info(&c)
	s := square{
		length:  4,
		breadth: 3,
	}
	info(&s)
	// info(s) will not work
}
