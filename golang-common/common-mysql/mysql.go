package main

import (
	"fmt"
	mysql "github.com/xiaozhiliaoo/golang-common/common-mysql"
)

func main() {
	db := mysql.GetDB()
	fmt.Printf("%s", db.DriverName())
}
