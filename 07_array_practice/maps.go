package main

import "fmt"

func main() {
	website := map[string]string{
		"github.com": "https://github.com",
		"golang.org": "https://golang.org",
	}
	fmt.Println(website)
	fmt.Println(website["golang.org"])
}
