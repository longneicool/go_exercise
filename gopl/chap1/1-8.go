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
		url = addPrefix(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, copyerr := io.Copy(os.Stdout, resp.Body)
		if copyerr != nil {
			fmt.Fprintf(os.Stderr, "copy: %v\n", err)
			os.Exit(1)
		}
	}
}

func addPrefix(url string) string {
	const httpPrefix = "http://"
	if !strings.HasPrefix(url, httpPrefix) {
		fmt.Println("The " + url + " doesn't have an prefix")
		return httpPrefix + url
	}

	return url
}
