package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

type BenchmarkTest02652 struct {
}

func (b *BenchmarkTest02652) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := r.URL.RawQuery
	paramVal := "BenchmarkTest02652="
	paramLoc := -1

	if queryString != "" {
		paramLoc = strings.Index(queryString, paramVal)
	}

	if paramLoc == -1 {
		http.Error(w, fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02652"), http.StatusBadRequest)
		return
	}

	param := queryString[paramLoc+len(paramVal):]
	if ampersandLoc := strings.Index(queryString[paramLoc:], "&"); ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : paramLoc+ampersandLoc]
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/goseccode")
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func doSomething(param string) string {
	a25323 := param
	b25323 := strings.Builder{}
	b25323.WriteString(a25323)
	b25323.WriteString(" SafeStuff")
	e25323 := b25323.String()[:len(b25323.String())-1]

	return e25323
}
