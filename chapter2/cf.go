// Cf converts its numberic argument to Celsius adn Fahrenheit

package main

import (
	"fmt"
	"hello/chapter2/tempconv"
	"os"
	"strconv"
)

func main() {

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cd: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FTocC(f),
			c, tempconv.CToF(c))
	}
}
