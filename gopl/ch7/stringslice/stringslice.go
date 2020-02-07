package main

import (
	"fmt"
	"sort"
)

//给切片加类型
type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	names := []string{"2", "4", "1", "100", "6"}
	sort.Sort(StringSlice(names))
	fmt.Println(names)

	names2 := []string{"2", "6", "4", "1", "5", "9", "3"}
	sort.Strings(names2)
	fmt.Println(names2)
}
