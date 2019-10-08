package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func (u user) AgeString() {
	fmt.Printf("%d : %s %s\n", u.Age, u.First, u.Last)
}

func (u user) LastString() {
	fmt.Printf("%s : %s\n", u.Last, u.First)
}

// ByAge type for custom sort based on Age
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByLast type for custom sort based on Last
type ByLast []user

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	fmt.Println("Users...")
	for _, v := range users {
		fmt.Printf("%s %s who is %d years old says:\n", v.First, v.Last, v.Age)
		for _, s := range v.Sayings {
			fmt.Printf("\t%s\n", s)
		}
	}
	fmt.Println("Sort Sayings for each user")
	for _, v := range users {
		sort.Strings(v.Sayings)
	}
	for _, v := range users {
		fmt.Printf("%s %s who is %d years old says:\n", v.First, v.Last, v.Age)
		for _, s := range v.Sayings {
			fmt.Printf("\t%s\n", s)
		}
	}

	sort.Sort(ByAge(users))
	for _, v := range users {
		v.AgeString()
	}

	sort.Sort(ByLast(users))
	for _, v := range users {
		v.LastString()
	}

}
