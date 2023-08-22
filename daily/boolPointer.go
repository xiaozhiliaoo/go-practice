package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type ChatRsp struct {
	Irrelevant *bool `json:"irrelevant"`
}

func main() {

	str := ""
	fmt.Println(len(str))

	rsp := ChatRsp{}
	fmt.Printf("%+v", rsp.Irrelevant)

	if true {
		isTrue := true
		rsp.Irrelevant = &isTrue
	}

	b, _ := jsoniter.Marshal(rsp)

	fmt.Printf("%+v", string(b))
}
