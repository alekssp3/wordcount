package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("test")
	fmt.Println(len(os.Args) - 1)
}
