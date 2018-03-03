package main

import (
			"fmt"
			"os"
			)

func main() {
	sep := " "
	for index, arg := range(os.Args[0:]) {
     fmt.Println(index, sep, arg)
	}
}