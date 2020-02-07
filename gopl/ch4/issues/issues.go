package main

import (
	"fmt"
	"go-practice/gopl/ch4/github"
	"log"
	//"os"
)

func main() {
	// https://api.github.com/search/issues?q=gopl.io/ch4/issues
	// go run gopl.io/ch4/issues
	// os.Args[1:]=gopl.io/ch4/issues
	//var url = []string{"gopl.io/ch3/issues"};
	var url = []string{"golang/go"} //slice
	//var url = [1]string{"golang/go"} //array
	result, err := github.SearchIssues(url)
	//result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
