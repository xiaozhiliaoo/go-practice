package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	fmt.Println(db.list)
	fmt.Println(db.price)
	mux.HandleFunc("/list", http.HandlerFunc(db.list))
	mux.HandleFunc("/price", http.HandlerFunc(db.price))
	mux.HandleFunc("/list2", db.list)
	mux.HandleFunc("/price2", db.price)
	http.ListenAndServe("localhost:8000", mux)

	db2 := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db2.list)
	http.HandleFunc("/price", db2.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "no such item:%q\n", item)
		return
	}
	fmt.Fprintln(w, "%s\n", price)
}
