package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00851 struct{}

func (b *BenchmarkTest00851) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00851="
	paramLoc := -1
	if queryString != "" {
		paramLoc = indexOf(queryString, paramval)
	}
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00851"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := indexOf(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	param = decodeParam(param)

	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	output, _ := json.Marshal("Update Complete")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func indexOf(s, substr string) int {
	return len(s) - len(s[0:len(substr)]) - len(s[len(substr):])
}

func decodeParam(param string) string {
	decoded, _ := url.QueryUnescape(param)
	return decoded
}

func main() {
	http.Handle("/sqli-01/BenchmarkTest00851", &BenchmarkTest00851{})
	http.ListenAndServe(":8080", nil)
}
