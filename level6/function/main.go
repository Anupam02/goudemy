package main

import "fmt"

func main() {
	foo()
	bar("James")
	r := woo("Moneypenny")
	fmt.Println(r)
	x, y := mouse("Ian", "Flemming")
	fmt.Println(x, y)
}

func foo() {
	fmt.Println("Hello from foo.")
}

func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo ", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false
	return a, b
}
