package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	users := []user{
		{"James", 32},
		{"Moneypenny", 27},
	}
	fmt.Println(users)
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
