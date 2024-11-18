package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func main() {
	http.HandleFunc("/sqli-02/BenchmarkTest01000", benchmarkTest01000)
	http.ListenAndServe(":8080", nil)
}

func benchmarkTest01000(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest01000",
			Value:  "verifyUserPassword%28%27foo%27%2C%27bar%27%29",
			MaxAge: 60 * 3,
			Secure: true,
			Path:   r.RequestURI,
			Domain: r.Host,
		})
		http.ServeFile(w, r, "sqli-02/BenchmarkTest01000.html")
		return
	}

	r.ParseForm()
	cookie, err := r.Cookie("BenchmarkTest01000")
	if err != nil {
		cookie = &http.Cookie{Name: "BenchmarkTest01000", Value: "noCookieValueSupplied"}
	}

	param, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		param = "noCookieValueSupplied"
	}

	bar := doSomething(param)
	sqlStr := fmt.Sprintf("CALL %s", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []map[string]interface{}
	columns, err := rows.Columns()
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		values := make([]sql.RawBytes, len(columns))
		scanArgs := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		if err := rows.Scan(scanArgs...); err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			rowMap[col] = string(values[i])
		}
		results = append(results, rowMap)
	}

	output, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func doSomething(param string) string {
	if param == "noCookieValueSupplied" {
		return ""
	}

	return string(param) // Replace with actual logic as needed
}
