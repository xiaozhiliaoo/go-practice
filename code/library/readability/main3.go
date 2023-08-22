package main

import (
	"fmt"
	"github.com/go-shiori/go-readability"
	"log"
	"time"
)

func main() {

	url := "https://www.beijing.gov.cn/"

	article, err := readability.FromURL(url, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to parse %s, %v\n", url, err)
	}

	fmt.Printf("URL     : %s\n", url)
	fmt.Printf("Title   : %s\n", article.Title)
	fmt.Printf("Author  : %s\n", article.Byline)
	fmt.Printf("Node  : %+v\n", *article.Node)
	fmt.Printf("Length  : %d\n", article.Length)
	fmt.Printf("Excerpt : %s\n", article.Excerpt)
	fmt.Printf("SiteName: %s\n", article.SiteName)
	fmt.Printf("Image   : %s\n", article.Image)
	fmt.Printf("Favicon : %s\n", article.Favicon)
	fmt.Printf("Language : %s\n", article.Language)
	fmt.Printf("Content : %s\n", article.Content)
	fmt.Printf("TextContent : %s\n", article.TextContent)
}
