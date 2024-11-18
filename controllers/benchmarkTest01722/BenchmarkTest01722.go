package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01722 struct{}

func (b *BenchmarkTest01722) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest01722="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParamIndex(queryString, paramval)
	}
	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01722"), http.StatusBadRequest)
		return
	}

	param := extractParam(queryString, paramLoc, paramval)
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	results, err := executeQuery(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Your results are: <br>"))
	for _, s := range results {
		w.Write([]byte(encodeForHTML(s) + "<br>"))
	}
}

func findParamIndex(queryString, paramval string) int {
	return -1 // Реализуйте логику поиска индекса параметра
}

func extractParam(queryString string, paramLoc int, paramval string) string {
	return "" // Реализуйте логику извлечения параметра
}

func doSomething(param string) string {
	return param // Здесь вы можете внести дополнительные изменения, если требуется
}

func executeQuery(sqlStr string) ([]string, error) {
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
	return results, nil
}

func encodeForHTML(input string) string {
	return input // Реализуйте логику экранирования для HTML
}

func main() {
	http.Handle("/sqli-03/BenchmarkTest01722", &BenchmarkTest01722{})
	http.ListenAndServe(":8080", nil)
}
