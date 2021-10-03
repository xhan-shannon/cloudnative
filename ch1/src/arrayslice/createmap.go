package main

import "fmt"

func printChar(ch string, iterNum int) string {
	var line string = ""
	for i := 0; i < iterNum; i++ {
		line = fmt.Sprintf("%s%s", line, ch)
	}
	return line
}

// func deleteItem(slice []int, index int) []int {
// 	return append(slice[:index], slice[index+1]...)
// }

func main() {
	aMap := map[string]func() int{
		"funcA": func() int {
			return 100
		},
	}

	fmt.Println("aMap is : ", aMap)
	f, exist := aMap["funcA"]
	if exist {
		fmt.Println("key existed and func call value: ", f())
	} else {
		fmt.Println("the key does not exist")
	}

	for k, v := range aMap {
		fmt.Println(k, v)
	}
}
