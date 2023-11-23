package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	s := `"1","2","3","4"`
	s = "[" + s + "]"
	fmt.Println(s)
	var ssss []string
	err := jsoniter.UnmarshalFromString(s, &ssss)
	fmt.Println(err)
	fmt.Println(ssss)

}
