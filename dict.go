package main

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "net/http"
    "log"
)

var global_db *sql.DB

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, getJson(r.URL.Path[1:]))
}

func getJson(word string) string {
	rows, err := global_db.Query("select definition from dictionary where word = ?", word)
	if err != nil {
		log.Fatal(err)
		return "{}"
	}
	defer rows.Close()
	for rows.Next() {
		var definition string
		rows.Scan(&definition)
		return fmt.Sprintf("{\"%s\": \"%s\"}", word, definition)
	}
	rows.Close()
	return "{}"
}

func main() {
	global_db, _ = sql.Open("sqlite3", "./dict.db")
	defer global_db.Close()

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}