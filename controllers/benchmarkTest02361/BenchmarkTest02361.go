package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02361Controller struct{}

func (c *BenchmarkTest02361Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := ""
	flag := true

	for name, values := range r.Form {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02361" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT TOP 1 USERNAME FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var result string
	err = db.QueryRow(sqlStr).Scan(&result)
	if err == sql.ErrNoRows {
		fmt.Fprintf(w, "No results returned for query: %s", sqlStr)
		return
	} else if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func doSomething(param string) string {
	bar := ""

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
