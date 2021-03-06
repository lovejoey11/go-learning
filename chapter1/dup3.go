package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		fmt.Println(filename) // exercise 1.4
		for line, n := range counts {
			if n > 1 {
				// fmt.Println(filename)
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}
