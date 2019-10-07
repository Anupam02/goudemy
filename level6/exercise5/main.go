package main

import "fmt"

type SQUARE struct {
	length  float64
	breadth float64
}

type CIRCLE struct {
	radius float64
}

func (s SQUARE) AREA() float64 {
	return s.length * s.breadth
}

func (c CIRCLE) AREA() float64 {
	return 3.14 * c.radius * c.radius
}

type SHAPE interface {
	AREA() float64
}

func INFO(s SHAPE) {
	a := s.AREA()
	fmt.Println(a)
}

func main() {
	s := SQUARE{
		length:  4,
		breadth: 5,
	}
	c := CIRCLE{
		radius: 4,
	}
	INFO(s)
	INFO(c)
}
