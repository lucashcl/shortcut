package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	urlRegex = `^https?://(www\.)?([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)(\.[a-zA-Z0-9]+)?(/.*)?$`
)

func main() {
	args := os.Args

	switch len(args) {
	case 1:
		var input string
		for {
			fmt.Print("Enter url: ")
			_, err := fmt.Scan(&input)
			if err != nil {
				fmt.Println(err)
				return
			}
			if isValidUrl(input) {
				break
			} else {
				fmt.Println("Invalid url")
			}
		}

		saveHtml(input)
	case 2:
		if isValidUrl(args[1]) {
			saveHtml(args[1])
		} else {
			fmt.Println("Invalid url")
		}
	}
}

func saveHtml(url string) {
	html := fmt.Sprintf(`<meta http-equiv="refresh" content="0; url=%s" />`, url)
	name := extractDomain(url)
	os.WriteFile(name+".html", []byte(html), 0644)
}

func extractDomain(url string) string {
	r, _ := regexp.Compile(urlRegex)
	matches := r.FindStringSubmatch(url)
	return strings.Join(matches[2:3], ".")
}

func isValidUrl(url string) bool {
	r, _ := regexp.Compile(urlRegex)
	return r.MatchString(url)
}
