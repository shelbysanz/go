package main

import (
	"fmt"
	"net/url"
)

func getDomainNameFromURL(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return parsedUrl.Hostname(), nil
}

func main() {
	// boot.dev
	fmt.Println(getDomainNameFromURL("https://boot.dev/learn/learn-python"))

	// youtube
	fmt.Println(getDomainNameFromURL("https://youtube.com"))
}
