package main

import (
	"fmt"
	"github.com/spf13/cast"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

const (
	getQAUnconfirmedNumByTime = `
		SELECT 
			count(*) 
		FROM 
		    t_doc_qa 
		WHERE 
		   corp_id = ? AND  robot_id = ? AND is_deleted = ? AND accept_status = ? AND create_time >= ?
	`
)

func main() {
	// 连接数据库
	db, err := sqlx.Open("mysql", "root:238888@tcp(localhost:3306)/db_task_interface")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	accessTime := time.UnixMilli(cast.ToInt64("1699436754099"))
	var total uint64
	args := []any{63, 419, 1, 2, accessTime}
	if err = db.Get(&total, getQAUnconfirmedNumByTime, args...); err != nil {

		fmt.Printf("err:%+v\n", err)
		return
	}
	fmt.Printf("count:%d", total)
}
