package main

import (
	crand "crypto/rand"
	"fmt"
	mrand "math/rand"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq" // enable support for Postgres
)

func main() {
	fmt.Println(mrand.Float64())
	fmt.Println(crand.Read([]byte{'1'}))

	//db, err = sql.Open("postgres", dbname) // OK
	//db, err = sql.Open("mysql", dbname) // OK
	//db, err = sql.Open("sqlite3", dbname)
}
