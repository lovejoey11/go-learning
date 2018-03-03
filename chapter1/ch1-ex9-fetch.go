package main

import (
	"fmt"
	// "io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("http status code: %s", resp.Status)
		// _, errc := io.Copy(os.Stdout, resp.Status)
		// resp.Body.Close()
		// if errc != nil {
		// 	fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)

		// }
		// fmt.Printf("%s", os.Stdout)
	}
}
