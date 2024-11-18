package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01719 struct{}

func (b *BenchmarkTest01719) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("BenchmarkTest01719")

	if param == "" {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01719"), http.StatusBadRequest)
		return
	}

	bar := b.doSomething(param)

	sqlStr := "SELECT * from USERS where USERNAME=? and PASSWORD='" + bar + "'"

	conn, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	stmt, err := conn.Prepare(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec("foo")
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	// Dummy response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func (b *BenchmarkTest01719) doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/", &BenchmarkTest01719{})
	http.ListenAndServe(":8080", nil)
}
