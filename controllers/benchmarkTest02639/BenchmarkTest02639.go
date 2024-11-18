package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02639 struct{}

func (b *BenchmarkTest02639) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02639="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParamLocation(queryString, paramval)
	}
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02639"), http.StatusBadRequest)
		return
	}

	param := extractParamValue(queryString, paramval, paramLoc)

	bar := doSomething(param)

	sqlQuery := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	err := executeSQLQuery(sqlQuery)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("No results can be displayed for query: %s<br> because the execute method doesn't return results.", sqlQuery)))
}

func findParamLocation(queryString, paramval string) int {
	return -1 // реализация поиска параметра
}

func extractParamValue(queryString, paramval string, paramLoc int) string {
	return "" // реализация извлечения значения параметра
}

func doSomething(param string) string {
	return param + " SafeStuff"
}

func executeSQLQuery(query string) error {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(query)
	return err
}
