package main

import (
	"fmt"
	"log"
	"os"
	"time"

	readability "github.com/go-shiori/go-readability"
)

var (
	urls = []string{
		// this one is article, so it's parse-able
		"https://www.nytimes.com/2019/02/20/climate/climate-national-security-threat.html",
		// while this one is not an article, so readability will fail to parse.
		"https://www.nytimes.com/",
	}
)

func main() {
	for i, url := range urls {
		article, err := readability.FromURL(url, 30*time.Second)
		if err != nil {
			log.Fatalf("failed to parse %s, %v\n", url, err)
		}

		dstTxtFile, _ := os.Create(fmt.Sprintf("text-%02d.txt", i+1))
		defer dstTxtFile.Close()
		dstTxtFile.WriteString(article.TextContent)

		dstHTMLFile, _ := os.Create(fmt.Sprintf("html-%02d.html", i+1))
		defer dstHTMLFile.Close()
		dstHTMLFile.WriteString(article.Content)

		fmt.Printf("URL     : %s\n", url)
		fmt.Printf("Title   : %s\n", article.Title)
		fmt.Printf("Author  : %s\n", article.Byline)
		fmt.Printf("Length  : %d\n", article.Length)
		fmt.Printf("Excerpt : %s\n", article.Excerpt)
		fmt.Printf("SiteName: %s\n", article.SiteName)
		fmt.Printf("Image   : %s\n", article.Image)
		fmt.Printf("Favicon : %s\n", article.Favicon)
		fmt.Printf("Text content saved to \"text-%02d.txt\"\n", i+1)
		fmt.Printf("HTML content saved to \"html-%02d.html\"\n", i+1)
		fmt.Println()
	}
}
