package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  string
	Addr []*Address
}

type Address struct {
	Location string
}

func main() {
	user := User{
		Name: "11",
		Age:  "22",
		Addr: []*Address{
			{
				Location: "33",
			},
			{
				Location: "44",
			},
		},
	}
	marshal, _ := json.Marshal(user)
	fmt.Printf("%+v\n", string(marshal))

	var newUser User
	json.Unmarshal(marshal, &newUser)

	fmt.Printf("%+v\n", newUser)

	location := newUser.Addr[0].Location

	fmt.Println(location)

}
