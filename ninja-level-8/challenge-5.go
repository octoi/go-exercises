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

type ByAge []user

func (x ByAge) Len() int { return len(x) }
func (x ByAge) Swap(i, j int) {x[i], x[j] = x[j], x[i]}
func (x ByAge) Less(i, j int) bool { return x[i].Age < x[j].Age }

type ByLast []user

func (x ByLast) Len() int { return len(x) }
func (x ByLast) Swap(i, j int) {x[i], x[j] = x[j], x[i]}
func (x ByLast) Less(i, j int) bool { return x[i].Last < x[j].Last }

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

	betterPrint(users, "default data")

	sort.Sort(ByAge(users))
	betterPrint(users, "sorted on age")

	sort.Sort(ByLast(users))
	betterPrint(users, "sorted on last")


	fmt.Println("\n=== sayings are sorted ===\n", "")
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
}

func betterPrint(users []user, s string) {
	fmt.Println("\n===", s, "===\n", "")

	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}
}
