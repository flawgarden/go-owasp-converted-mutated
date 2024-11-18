package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type BenchmarkTest02283 struct {
	Db *sql.DB
}

func NewBenchmarkTest02283(db *sql.DB) *BenchmarkTest02283 {
	return &BenchmarkTest02283{Db: db}
}

func (b *BenchmarkTest02283) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Query().Get("BenchmarkTest02283")
	param := ""
	if vals, ok := r.URL.Query()["BenchmarkTest02283"]; ok && len(vals) > 0 {
		param = vals[0]
	}

	bar := b.doSomething(r, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	statement, err := b.Db.Prepare(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	results, err := statement.Query()
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer results.Close()

	// Process results...
}

func (b *BenchmarkTest02283) doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[1:], valuesList[2]) // remove the 1st safe value
		bar = valuesList[1]                                // get the last 'safe' value
	}
	return bar
}
