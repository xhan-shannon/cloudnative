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
	var fullString string = "hello world!"

	for _, ch := range fullString {
		fmt.Printf("%s\n", string(ch))
	}
}
