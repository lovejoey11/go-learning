package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToString([]int{1, 2, 3, 4}))
	fmt.Println(comma_nr("1332215610162144"))
	fmt.Println(comma_fl("1332215610162144.003974"))
	fmt.Println(anagram("ananananana", "ananananana"))
}

func intToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func comma_nr(s string) string {
	var buf bytes.Buffer
	if len(s) <= 3 {
		fmt.Fprintf(&buf, "%s", s)
	}
	num := 0
	for i := len(s) - 1; i >= 0; i-- {
		if num == 3 {
			buf.WriteString(",")
			num = 0
		}
		fmt.Fprintf(&buf, "%c", s[i])
		num++
	}
	rString := buf.String()
	buf.Reset()
	for i := len(rString) - 1; i >= 0; i-- {
		fmt.Fprintf(&buf, "%c", rString[i])
	}
	return buf.String()
}

func comma_fl(s string) string {
	var buf bytes.Buffer
	if strings.Contains(s, ".") {
		dotindex := strings.Index(s, ".")
		fmt.Fprintf(&buf, "%s", comma_nr(s[:dotindex]))
		buf.WriteString(s[dotindex:])
		// fmt.Println(dotindex, len(s))
		return buf.String()
	} else {
		return comma_nr(s)
	}
}
func anagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var buf bytes.Buffer
	for i := len(s1) - 1; i >= 0; i-- {
		fmt.Fprintf(&buf, "%c", s1[i])
	}
	if strings.Compare(s2, buf.String()) == 0 {
		return true
	} else {
		return false
	}
}
