package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(splitString("123456488", 2))

	fmt.Println(len(`1这是普通文本2这是超链接,<a href="http">百度</a>`))
	fmt.Println(len("啊李"))
	str := []string{"123", "456", "789"}
	strings := merge(str, 6)
	fmt.Printf("%s", strings)
}

func ddd(res []string, maxByteSize int) []string {
	var currentString string = ""
	var targetString []string

	for _, str := range res {
		if len(currentString)+len(str) > maxByteSize {
			targetString = append(targetString, currentString)
			currentString = ""
		}
		currentString += str // 累积字符串
	}

	if currentString != "" {
		targetString = append(targetString, currentString)
	}

	return targetString
}

func merge(text []string, maxByteSize int) []string {

	currentString := ""
	var rst []string

	for _, str := range text {
		if len(currentString)+len(str) >= maxByteSize {
			rst = append(rst, currentString)
			currentString = ""
		}
		currentString += str
	}

	if currentString != "" {
		rst = append(rst, currentString)
	}

	return rst
}

func splitString(str string, chunkSize int) []string {
	strLength := len(str)
	chunks := int(math.Ceil(float64(strLength) / float64(chunkSize)))
	result := make([]string, chunks)

	start, end := 0, 0
	for i := 0; i < chunks; i++ {
		start = i * chunkSize
		end = (i + 1) * chunkSize
		if end > strLength {
			end = strLength
		}
		result[i] = str[start:end]
	}

	return result
}
