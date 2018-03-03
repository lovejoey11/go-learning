package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "foo/bar.go"
	b := "a.b.c.go"
	c := "12345"
	d := "234"
	fmt.Printf("%s\n", basename(a))
	fmt.Printf("%s\n", basename(b))

	fmt.Printf("%s\n", basename2(a))
	fmt.Printf("%s\n", basename2(b))

	fmt.Printf("%s\n", comma(c))
	fmt.Printf("%s", comma(d))

}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]

}
func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s

}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
