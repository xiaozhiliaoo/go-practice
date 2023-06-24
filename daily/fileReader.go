package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		panic(err)
	}

	var files []string

	err = filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, f := range files {
		content, err := ioutil.ReadFile(f)
		if err != nil {
			panic(err)
		}

		fmt.Printf("读取到的 YAML 内容：%s\n", content)
	}
}
