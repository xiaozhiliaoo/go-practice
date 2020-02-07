package main

import (
	"fmt"
	"os"
)

func tempDirs2() []string {
	return []string{"ddd", "dddww", "ece"}
}

func main() {
	var rmdirs []func()
	dirs := tempDirs2()
	for i := 0; i < len(dirs); i++ {
		os.MkdirAll(dirs[i], 0755) // OK
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dirs[i]) // NOTE: incorrect!
		})
		fmt.Println(rmdirs)
	}

}
