package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	truck1 := truck{
		vehicle{2,
			"black",
		},
		true,
	}
	fmt.Println(truck1)
	fmt.Println(truck1.doors)
	sedan1 := sedan{
		vehicle{
			4,
			"black",
		},
		true,
	}
	fmt.Println(sedan1)
	fmt.Println(sedan1.doors)
}
