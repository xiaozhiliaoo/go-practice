package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m)
	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1" (first value)
	fmt.Println(m["item"])
	m = nil
	m.Get("item")
	m.Add("item", "3")
}
