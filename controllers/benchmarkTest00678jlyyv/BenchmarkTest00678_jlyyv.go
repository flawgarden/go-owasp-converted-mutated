package controllers

import (
"database/sql"
"encoding/json"
"fmt"
"net/http"
_ "github.com/go-sql-driver/mysql"
"sync"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00678 struct{}

func (b *BenchmarkTest00678) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00678")
	if param == "" {
		param = ""
	}

	var bar string
	num := 196
	if (500/42)+num > 200 {

w := &Wrapper[string]{Value: param}
task1 := NewSettingTask(w, "")
task2 := NewSettingTask(w, param)
var wg sync.WaitGroup
wg.Add(2) // Добавляем 2 задачи в WaitGroup
go func() {
    defer wg.Done()
    task1.Run()
}()
go func() {
    defer wg.Done()
    task2.Run()
}()
wg.Wait()
param = w.Value

		bar = param
	} else {
		bar = "This should never happen"
	}

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var results int64
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		http.Error(w, "No results returned for query: "+sqlStr, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"results": results}
	output, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
