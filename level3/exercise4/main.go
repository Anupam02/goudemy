package main

import "fmt"

func main() {
	yborn := 1993
	ycurr := 2019
	i := yborn
	for {
		if i > ycurr {
			break
		}
		fmt.Printf("%v\t", i)
		i++
	}
}
