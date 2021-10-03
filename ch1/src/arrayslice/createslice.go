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
	var slice1 []int
	slice1 = make([]int, 0)
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	fmt.Printf("slice1: %+v\n", slice1)

	a := []int
	b := []int{1, 2, 3}
	c := a
	a = append(b, 1)
	fmt.Println("c is : ", c)

}
