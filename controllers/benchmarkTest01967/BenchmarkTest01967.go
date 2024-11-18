package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type BenchmarkTest01967 struct{}

func (b *BenchmarkTest01967) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest01967")
	sqlStr := fmt.Sprintf("SELECT USERNAME FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", param)

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/goseccode")
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var result string
	err = db.QueryRow(sqlStr).Scan(&result)
	if err != nil {
		http.Error(w, "No results returned for query", http.StatusNotFound)
		return
	}

	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func main() {
	http.Handle("/sqli-04/BenchmarkTest01967", &BenchmarkTest01967{})
	http.ListenAndServe(":8080", nil)
}
