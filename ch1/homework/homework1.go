package main

import "fmt"

func main() {
	var sarr []string = []string{
		"I",
		"am",
		"stupid",
		"and",
		"weak",
	}

	for index, _ := range sarr {
		if index == 2 {
			sarr[index] = "smart"
		} else if index == 4 {
			sarr[index] = "strong"
		}
	}

	fmt.Println(sarr)
}
