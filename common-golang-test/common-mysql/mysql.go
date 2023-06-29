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

func path() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	exePath := os.Args[0]
	return filepath.Join(dir, exePath)
}

func main() {
	dir, _ := os.Getwd()
	fmt.Printf("dir:%s\n", dir)
	mysql.Init(dir + "/common-mysql/config.yaml")

	db := mysql.GetDB()
	fmt.Printf("%s", db.DriverName())
}
