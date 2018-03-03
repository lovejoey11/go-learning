// Cf converts its numberic argument to Celsius adn Fahrenheit

package main

import (
	"fmt"
	"hello/chapter2/commconv"
	// "hello/chapter2/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[2:] {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cd: %v\n", err)
			os.Exit(1)
		}

		switch os.Args[1] {
		case "weight":
			weiconv(value)
		case "length":
			lengthconv(value)
		default:
			fmt.Printf("unexpected command: %v", os.Args[1])
		}
	}
}

func weiconv(value float64) {
	lb := commconv.Pound(value)
	fmt.Printf("%s = %s\n", lb, commconv.PoundToKg(lb))
}

func lengthconv(value float64) {
	cm := commconv.Centimeter(value)
	fmt.Printf("%s = %s\n",
		cm, commconv.CentiToInch(cm))
}
