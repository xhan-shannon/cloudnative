package main

import "fmt"

func printChar(ch string, iterNum int) string {
	var line string = ""
	for i := 0; i < iterNum; i++ {
		line = fmt.Sprintf("%s%s", line, ch)
	}
	return line
}
func main() {
	fmt.Println("in forloop")
	var diagLevel = 10
	for i := 1; i < 10; i++ {
		var line string
		line = printChar(" ", diagLevel-i)
		line += printChar("*", i)
		fmt.Println(line)
	}
}
