//submitted by Rocket dev team:
// Liz Aharonian - 316584960
// Raz Shenkman - 311130777
// Ori Ben Zaken - 311492110

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//change url according to req
		url = addPrefix(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//print status code
		fmt.Fprintf(os.Stdout, "resp.status %s\n", resp.Status)
		//io.Copy uses fixed 32KB buffer to copy from reader to writer until EOF.
		//So no matter how big the source is, you’ll always just use 32KB to copy it to destination.
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func addPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	return url
}
