package main

import "fmt"

func main() {
	favSport := "TT"
	switch favSport {
	case "VolleyBall":
		fmt.Println("Favourite Game is VolleyBall")
	case "FootBall":
		fmt.Println("Favourite Game is FootBall")
	case "TT":
		fmt.Println("Favourite Came is TT")
	default:
		fmt.Println("default")
	}
}
