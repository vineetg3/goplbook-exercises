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
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", err)
			continue
		}
		bytesRead, error := io.Copy(os.Stdout, resp.Body)
		fmt.Println("---------------")
		fmt.Printf("Response code: %s\n", resp.Status)
		fmt.Printf("Bytes read: %d\n", bytesRead)
		if error != nil {
			fmt.Fprintf(os.Stderr, "Error: %v", error)
			continue
		}
		fmt.Println("---------------")
		fmt.Println("---------------")
	}
}
