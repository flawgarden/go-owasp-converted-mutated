package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02732 struct{}

func (b *BenchmarkTest02732) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest02732) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02732")
	bar := doSomething(r, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	results, err := queryDatabase(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Your results are: <br>"))
	for _, s := range results {
		w.Write([]byte(s + "<br>"))
	}
}

func doSomething(r *http.Request, param string) string {
	return param
}

func queryDatabase(sqlStr string) ([]string, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		results = append(results, username)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func main() {
	http.Handle("/sqli-06/BenchmarkTest02732", &BenchmarkTest02732{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
		os.Exit(1)
	}
}
