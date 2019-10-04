package main

import "fmt"

func main() {
	fav := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	for k, v := range fav {
		fmt.Printf("%v likes : \n", k)
		for i, f := range v {
			fmt.Printf("\t%d\t%v\n", i, f)
		}
	}
}
