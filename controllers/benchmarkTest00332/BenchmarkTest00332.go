package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

type BenchmarkTest00332 struct {
	DB *sql.DB
}

func (b *BenchmarkTest00332) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest00332) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.Header.Get("BenchmarkTest00332")
	param, _ = url.QueryUnescape(param)

	a40477 := param
	b40477 := a40477 + " SafeStuff"
	b40477 = b40477[:len(b40477)-len("Chars")] + "Chars"

	map40477 := map[string]interface{}{
		"key40477": b40477,
	}
	c40477 := map40477["key40477"].(string)
	d40477 := c40477[:len(c40477)-1]

	f40477 := string([]byte(d40477))
	bar := f40477 // Static input for safe flow

	sqlStr := "SELECT * from USERS where USERNAME=? and PASSWORD='" + bar + "'"

	stmt, err := b.DB.Prepare(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	var results []map[string]interface{}
	err = stmt.QueryRow("foo").Scan(&results)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/goseccode")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.Handle("/sqli-00/BenchmarkTest00332", &BenchmarkTest00332{DB: db})
	http.ListenAndServe(":8080", nil)
}
