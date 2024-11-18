package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02551 struct{}

func (b *BenchmarkTest02551) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest02551="
	paramLoc := -1
	if queryString != "" {
		paramLoc = findParamIndex(queryString, paramval)
	}

	if paramLoc == -1 {
		fmt.Fprintf(w, "getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02551")
		return
	}

	param := extractParamValue(queryString, paramLoc, paramval)
	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Write(output)
}

func findParamIndex(queryString, paramval string) int {
	return findIndex(queryString, paramval)
}

func extractParamValue(queryString string, paramLoc int, paramval string) string {
	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := findAmpersand(queryString, paramLoc)
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : ampersandLoc]
	}
	return param
}

func findAmpersand(queryString string, paramLoc int) int {
	return findIndex(queryString[paramLoc:], "&")
}

func findIndex(s, substr string) int {
	return -1 // Stub for the index finding logic
}

func doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		bar = param[:len(param)-1] + "Z"
	}
	return bar
}

func main() {
	http.Handle("/crypto-02/BenchmarkTest02551", &BenchmarkTest02551{})
	http.ListenAndServe(":8080", nil)
}
