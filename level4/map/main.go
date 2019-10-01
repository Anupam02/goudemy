package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Anupam"])

	v, ok := m["Anupam"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Anupam"]; ok {
		fmt.Println(v)
	}
	if v, ok := m["James"]; ok {
		fmt.Println(v)
	}
	m["todd"] = 33
	for k, v := range m {
		fmt.Println(k, v)
	}

	// xi := []int{3, 4, 5, 6, 7, 9}
	// for i, v := range xi {
	// 	fmt.Println(i, v)
	// }
	delete(m, "James")
	fmt.Println(m)
	fmt.Println(m["James"])
	delete(m, "Doesn't exists")
	fmt.Println(m)
	fmt.Println(m["Doesn't exists"])

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}
	fmt.Println(m)
}
