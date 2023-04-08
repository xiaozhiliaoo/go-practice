package main

import (
	"fmt"
	"sort"
)

type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

func (s Sequence) String() string {
	s = s.Copy()
	sort.Sort(s)
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

func main() {
	var s Sequence = []int{7, 8, 9, 234, 999, 111, 1, 2, 3, 4, 5}
	sort.Sort(s)
	fmt.Println(s)
}
