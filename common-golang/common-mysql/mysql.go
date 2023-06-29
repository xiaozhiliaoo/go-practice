package main

import (
	"fmt"
	mysql "github.com/xiaozhiliaoo/common-golang/common-mysql"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {
	directory := getCurrentDirectory()
	fmt.Printf("%s", directory)
	mysql.Init(directory + "/config.yaml")
	db := mysql.GetDB()
	fmt.Printf("%s", db.DriverName())
}
