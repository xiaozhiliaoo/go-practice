package main

import (
	"os"
)

func main() {
	//函數切片
	var rmdirs []func()

	for _, dir := range tempDirs() {
		dir := dir             // NOTE: necessary!
		os.MkdirAll(dir, 0755) // creates parent directories too
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	for _, rmdir := range rmdirs {
		rmdir()
	}
}

func tempDirs() []string {
	return []string{"ddd", "dddww", "ece"}
}
