package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	fmt.Println(cast.ToUint8(true))
	fmt.Println(cast.ToUint8(false))
	link := &RefLink{
		"dd", "ddd",
	}
	fmt.Printf("%+v|%t", link, true)

}

func processRefLink(links []RefLink) {
	for i, item := range links {
		links[i].Link = addQueryString(item.Link)
	}
}

func addQueryString(link string) string {
	return link + "alibaba======="
}

type RefLink struct {
	Link string `json:"link"`
	Text string `json:"text"`
}
