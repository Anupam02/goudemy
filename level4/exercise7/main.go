package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken", "not stirred"}
	mm := []string{"Miss", "Moneypenny", "Helloooooo", "James"}
	x := [][]string{jb, mm}
	fmt.Println(x)
	for i := range x {
		for j := range x[i] {
			fmt.Println(x[i][j])
		}
	}
}
