package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00766Controller struct {
	http.Handler
}

func (c *BenchmarkTest00766Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest00766Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00766")
	if param == "" {
		param = ""
	}

	num := 106
	bar := ""
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var results int64
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		http.Error(w, "No results returned for query: "+sqlStr, http.StatusNotFound)
		return
	}
	w.Write([]byte("Your results are: " + strconv.FormatInt(results, 10)))
}
