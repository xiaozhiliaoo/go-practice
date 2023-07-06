package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"reflect"
	"time"
)

var db *sql.DB

type Course struct {
	Id          int64     `db:"id"`           // 主键id
	Name        string    `db:"name"`         // 课程名字
	PersonId    int64     `db:"person_id"`    //
	Open        int64     `db:"open"`         // 是否开课(0：开课 1：没开课)
	Teacher     string    `db:"teacher"`      // 老师名字
	GmtCreate   time.Time `db:"gmt_create"`   // 创建时间
	GmtModified time.Time `db:"gmt_modified"` // 修改时间
}

// closing bad idle connection: EOF
func main() {
	http.HandleFunc("/", selectQuery)
	http.HandleFunc("/sleep", lowSql)
	http.HandleFunc("/db_stats", dbStats)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func init() {
	db = getDB()

}

func selectQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "db.Stats():%+v\n", db.Stats())
	course := query()
	fmt.Fprintf(w, "course:%+v\n", course)
	fmt.Fprintf(w, "db.Stats():%+v\n", db.Stats())

}

func lowSql(w http.ResponseWriter, r *http.Request) {
	course := sleep()
	fmt.Fprintf(w, "course:%+v", course)
	printDBStat()
}

func dbStats(w http.ResponseWriter, r *http.Request) {
	stats := db.Stats()
	fmt.Fprintf(w, "course:%+v", stats)
}

func printDBStat() {
	stats := db.Stats()
	fmt.Printf("db stats:%+v", stats)
}

func query() Course {
	rows, err := db.Query("select * from course where id=1")
	if err != nil {
		fmt.Printf("db.Query err %+v", err)
	}
	defer rows.Close()
	var course Course
	for rows.Next() {

		s := reflect.ValueOf(&course).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(columns...)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(course)

	}
	return course
}

func sleep() Course {

	db.Query("SELECT SLEEP(30);")

	rows, _ := db.Query("select * from course where id=1")
	defer rows.Close()
	var course Course
	for rows.Next() {

		s := reflect.ValueOf(&course).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(columns...)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(course)

	}
	return course
}

func getDB() *sql.DB {
	db, err := sql.Open("mysql", "root:238888@tcp(127.0.0.1:3306)/training?charset=utf8mb4&timeout=1s&parseTime=true&interpolateParams=true&loc=Local")
	if err != nil {
		log.Printf("sql.Open err:%+v", err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(30 * time.Second)
	return db
}
