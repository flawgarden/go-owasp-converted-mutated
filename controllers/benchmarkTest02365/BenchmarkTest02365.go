package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02365 struct{}

func (b *BenchmarkTest02365) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	param := ""
	flag := true

	for name, values := range r.URL.Query() {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02365" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := doSomething(r, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

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

	// Handle response (if necessary)
	w.Header().Set("Content-Type", "application/json")
	output := map[string]string{"status": "success"}
	json.NewEncoder(w).Encode(output)
}

func doSomething(req *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[1]         // get the last 'safe' value
	}
	return bar
}
