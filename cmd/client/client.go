package main

import (
    "fmt"
    "strings"
)

type WebsiteURL struct {
	url1 string
	url2 string
    url3 string
}

func main() {
    url := WebsiteURL{"google.com", "youtube.com", "moodle.com"}
    join(&url)
}

func join(v *WebsiteURL) {
    urls := []string{v.url1, v.url2, v.url3}
    for _, u := range urls {
        fullURL := strings.Join([]string{"http://", u}, "")
        fmt.Println(fullURL)
    }
}