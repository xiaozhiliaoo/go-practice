package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
)

var str1 = `{"str":"1986-03-13T00:00:00.000","is_finish":true}`

type Data struct {
	IsFinish bool   `json:"is_finish"`
	Str      string `json:"str"`
}

type All struct {
	DataStr  string `json:"data_str"`
	DataStr2 string `json:"data_str2"`
}

func main() {

	var data []Data
	data = append(data, Data{IsFinish: true, Str: "str1"})
	data = append(data, Data{IsFinish: false, Str: "str2"})
	data = append(data, Data{IsFinish: false, Str: "str3"})
	toString, _ := jsoniter.MarshalToString(data)

	sprintf := fmt.Sprintf("%s", toString)

	toString2, _ := jsoniter.MarshalToString(sprintf)
	fmt.Println(toString2)

	all := All{DataStr: toString, DataStr2: "111"}
	fmt.Printf("%+v\n \n", all)

	marshalToString, _ := jsoniter.MarshalToString(toString)
	fmt.Println(marshalToString)

	var em Data
	err := jsoniter.UnmarshalFromString(str1, &em)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(em)

	aaa := []string{"1", "2", "3"}
	aaa2, _ := jsoniter.MarshalToString(aaa)
	fmt.Println(aaa2)
}
