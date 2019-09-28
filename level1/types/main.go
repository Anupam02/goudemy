package main

import "fmt"

var x bool

type myownbool bool

var y myownbool

var c int8 = 127

const p = 50
const q = 50.0001
const r = "James Bond"

const (
	m = 50
	n = 50.93883
	o = "Anupam Patel"
)

// typed constants
const (
	e int     = 234
	f float64 = 34.2342
	g string  = "Sahil Virmani"
)

func main() {
	fmt.Println(x, y)
	x = true
	y = true
	a := 42
	b := 42.34543
	s := "Hello There!!"
	bs := []byte(s)
	fmt.Println(x, y, a, b, s)
	fmt.Println(bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U \n", s[i])
	}
	fmt.Println("")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d => %x\n", i, s[i])
	}
	c := "H"
	fmt.Println(c)
	cs := []byte(c)
	fmt.Println(cs)

	n := cs[0]
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)
	fmt.Println(p, m, e)
	fmt.Println(q, n, f)
	fmt.Println(r, o, g)
	fmt.Printf("%T\t%T\t%T\n", p, m, e)
	fmt.Printf("%T\t%T\t%T\n", q, n, f)
	fmt.Printf("%T\t%T\t%T\n", r, o, g)
}
