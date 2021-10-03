package main

import "fmt"

func printChar(ch string, iterNum int) string {
	var line string = ""
	for i := 0; i < iterNum; i++ {
		line = fmt.Sprintf("%s%s", line, ch)
	}
	return line
}

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1]...)
}

func main() {
	arrayOne := [5]int{1, 2, 3, 4, 5}
	sliceOne := arrayOne[1:3]
	emptySlice := []int{}

	slice1 := append(emptySlice, 1)
	slice1 = append(slice1, 2)
	fmt.Printf("arrayOne: %+v, type: %T, address: %p\n", arrayOne, arrayOne, &arrayOne)
	fmt.Printf("sliceOne: %+v, type: %T, address: %p\n", sliceOne, sliceOne, sliceOne)
	fmt.Printf("emptySlice: %+v, type: %T, address: %p\n", emptySlice, emptySlice, emptySlice)
	fmt.Printf("slice1: %+v\n", slice1)

}
