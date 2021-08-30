package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var ans int
	if len(os.Args) > 1 && os.Args[1] != "" {
		ans = len(strings.Split(os.Args[1], " "))
	}
	fmt.Println(ans)
}
