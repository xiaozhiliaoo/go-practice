package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	uin := "700000920773"

	params := make(map[string]interface{})
	inter := make(map[string]interface{})
	uinParams := make(map[string]interface{})
	params["version"] = 1
	params["componentName"] = "MC"
	params["eventId"] = 1103941038
	params["interface"] = inter
	inter["interfaceName"] = "qcloud.Quser.getNickname"
	uinParams["uin"] = uin
	inter["para"] = uinParams

	// 将 map 转换为 JSON 格式
	jsonData, err := json.Marshal(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印 JSON 数据
	fmt.Println(string(jsonData))
}
