package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02184 struct{}

func (b *BenchmarkTest02184) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02184")
	if param == "" {
		param = ""
	}

	bar := doSomething(r, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	results, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer results.Close()

	w.Write([]byte("Your results are: "))
	for results.Next() {
		var username string
		if err := results.Scan(&username); err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(fmt.Sprintf("%s ", username)))
	}
}

func doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/sqli-04/BenchmarkTest02184", &BenchmarkTest02184{})
	http.ListenAndServe(":8080", nil)
}
