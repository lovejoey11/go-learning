package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	option := ""
	if len(os.Args) > 2 {
		option = os.Args[2]
	} else {
		option = "default"
	}
	vars := os.Args[1]
	switch option {
	case "sha384":
		fmt.Printf("SHA384 of %s is %x", vars, sha512.Sum384([]byte(vars)))
	case "sha512":
		fmt.Printf("SHA512 of %s is %x", vars, sha512.Sum512([]byte(vars)))
	default:
		fmt.Printf("SHA256 of %s is %x", vars, sha256.Sum256([]byte(vars)))

	}

}
