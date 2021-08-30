package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("0")
	} else {
		data := strings.Split(os.Args[1], " ")
		fmt.Println(len(data))
	}
}
