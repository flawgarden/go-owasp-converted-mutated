package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00767 struct {
}

func (b *BenchmarkTest00767) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values := r.Form["BenchmarkTest00767"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := ""
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C':
	case 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	sqlStr := fmt.Sprintf("SELECT userid from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var results int64

map787234 := make(map[string]string)
map787234["UbmbS"] = "JTjTv"
map787234["UbmbS"] = guess
sqlStr = map787234["UbmbS"]
map787234["UbmbS"] += guess

	err = db.QueryRow(sqlStr).Scan(&results)
	if err == sql.ErrNoRows {
		http.Error(w, "No results returned for query: "+sanitizeSQL(sqlStr), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func sanitizeSQL(sql string) string {
	return strings.ReplaceAll(sql, "<", "&lt;")
}
