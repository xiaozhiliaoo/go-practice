package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/go-shiori/go-readability"
)

var (
	urls2 = []string{
		"https://golangnote.com/topic/260.html",
	}
)

// https://golangnote.com/topic/305.html 正文识别算法
func main() {
	for _, u := range urls2 {
		resp, err := http.Get(u)
		if err != nil {
			log.Fatalf("failed to download %s: %v\n", u, err)
		}
		defer resp.Body.Close()

		ur, _ := url.Parse(u)
		article, err := readability.FromReader(resp.Body, ur)
		if err != nil {
			log.Fatalf("failed to parse %s: %v\n", u, err)
		}

		fmt.Printf("URL     : %s\n", u)
		fmt.Printf("Title   : %s\n", article.Title)
		fmt.Printf("Author  : %s\n", article.Byline)
		fmt.Printf("Length  : %d\n", article.Length)
		fmt.Printf("Excerpt : %s\n", article.Excerpt)
		fmt.Printf("SiteName: %s\n", article.SiteName)
		fmt.Printf("Image   : %s\n", article.Image)
		fmt.Printf("Favicon : %s\n", article.Favicon)
		fmt.Printf("Content : %s\n", article.Content)
		fmt.Printf("TextContent : %s\n", article.TextContent)
		fmt.Println()
	}
}
