package main

import "fmt"

func main() {
	m := map[string][]string {
		"bond_james": {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr": {`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["golang"] = []string{"Awesome", "Amazing", "cool"}

	for val, data := range m {
		fmt.Println("Record for ", val)

		for idx, item := range  data {
			fmt.Println("\t", idx, item)
		}
	}
}