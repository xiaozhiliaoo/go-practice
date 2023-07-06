package main

import (
	"context"
	"flag"
	"fmt"
	log "github.com/xiaozhiliaoo/common-golang/common-log"
	mysql "github.com/xiaozhiliaoo/common-golang/common-mysql"
	"os"
	"path/filepath"
	"time"
)

type Course struct {
	Id          int64     `db:"id"`           // 主键id
	Name        string    `db:"name"`         // 课程名字
	PersonId    int64     `db:"person_id"`    //
	Open        int64     `db:"open"`         // 是否开课(0：开课 1：没开课)
	Teacher     string    `db:"teacher"`      // 老师名字
	GmtCreate   time.Time `db:"gmt_create"`   // 创建时间
	GmtModified time.Time `db:"gmt_modified"` // 修改时间
}

func path() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	exePath := os.Args[0]
	return filepath.Join(dir, exePath)
}

// go build common-mysql/main.go
// ./main -conf /Users/jiandanlli/Desktop/gitspace/github/go-practice/common-golang-test/common-mysql/config.yaml
func main() {
	configFile := flag.String("conf", "", "application config file path")
	flag.Parse()

	dir, _ := os.Getwd()
	fmt.Printf("dir:%s\n", dir)
	mysql.Init(*configFile)
	log.Init(*configFile)

	log.InfoContextf(context.Background(), "dir is:%s", dir)
	db := mysql.GetDB()
	fmt.Printf("db drive name:%s", db.DriverName())

	name := db.DriverName()
	fmt.Printf("DriverName:%s", name)
	query, err := db.Query("select * from course")
	if err != nil {
		fmt.Printf("query err:%+v", err)
	}
	columns, _ := query.Columns()
	for _, column := range columns {
		fmt.Printf("col:%s\n", column)
	}

	course := []Course{}
	err = db.Select(&course, "SELECT * FROM course")
	if err != nil {
		fmt.Printf("query err:%+v", err)
	}
	for _, c := range course {
		fmt.Printf("course is:%+v\n", c)
	}
}
