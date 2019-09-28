package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}
	fmt.Println("")
	x := 1
	for x < 10 {
		fmt.Print(x)
		x++
	}
	fmt.Println("")
	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Print(y)
		y++
	}
	fmt.Println("")
	z := 1
	for {
		z++
		if z > 10 {
			break
		}
		if z%2 != 0 {
			continue
		}
		fmt.Print(z)
	}
}
