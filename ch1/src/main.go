package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	fmt.Println("os args: ", os.Args)
	fmt.Println(
		"the parameter name value is: ",
		*name,
	)
	fmt.Printf("os %s\n", name)
	fmt.Printf("%d\n", name, name)
}
