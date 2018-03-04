package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	diff := 0
	for i, _ := range c1 {
		if c1[i] != c2[i] {
			diff++
		}
	}
	fmt.Printf("The difference is %d", diff)
}
