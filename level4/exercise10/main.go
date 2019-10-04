package main

import "fmt"

func main() {
	fav := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	if _, ok := fav["patel_anupam"]; !ok {
		fav["patel_anupam"] = []string{"Golang", "Python", "Bash"}
	}
	fmt.Println(fav)
	if _, ok := fav["patel_anupam"]; ok {
		delete(fav, "patel_anupam")
	}
	fmt.Println(fav)

}
