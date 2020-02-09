package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		//当前目录
		roots = []string{"."}
		//roots := []string{"."}
	}

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			//fileSizes is send channel
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	//main goroutine 从 channel接收并输出
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f MB\n", nfiles, float64(nbytes)/1e6)
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}

	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "du1:%v\n", err)
	}
	return entries
}
