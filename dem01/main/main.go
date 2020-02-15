package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Bitcoin and ", os.Args[1])
	}
	fmt.Println("Hello World")
}
