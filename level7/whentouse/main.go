// everything in go is pass by value
package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x in main before calling foo", x)
	foo(x)
	fmt.Println("x in main after calling foo", x)
	y := 0
	fmt.Println("y in main before calling bar", y)
	bar(&y)
	fmt.Println("y in main after calling bar", y)

}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func bar(y *int) {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 43
	fmt.Println(y)
	fmt.Println(*y)
}
