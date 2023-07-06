package main

import (
	"fmt"
	"time"
)

var Formatter = "2006-01-02"

func main() {
	str2Time := convertStr2Time("2023-10-05")
	fmt.Println(str2Time)
}

func convertStr2Time(date string) time.Time {
	fromTime, _ := time.Parse(Formatter, date)
	return fromTime
}
