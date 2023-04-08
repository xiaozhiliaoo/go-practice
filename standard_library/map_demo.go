package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	nested := map[string]map[string]string{
		"1": {
			"a": "Apple",
			"b": "Banana",
			"c": "Coconut",
		},
	}
	fmt.Println("---start---")
	fmt.Println(nested["2"]["33"])
	_, ok := nested["2"]
	fmt.Println(ok)

	_, ok2 := nested["2"]["33"]
	fmt.Println(ok2)

	blank := DefaultIfBlank(nested["2"]["33"], "ddddd")
	fmt.Println(blank)

	blank2 := DefaultIfBlank(nested["1"]["a"], "ddddd")
	fmt.Println(blank2)

	fmt.Println("---end---")

	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)

	open, err := os.Open("22")

}

func DefaultIfBlank(str string, defaultStr string) string {
	if str != "" {
		return str
	}
	return defaultStr
}
