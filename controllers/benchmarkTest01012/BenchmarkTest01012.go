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

type BenchmarkTest01012 struct {
	DB *sql.DB
}

func NewBenchmarkTest01012() (*BenchmarkTest01012, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	return &BenchmarkTest01012{DB: db}, nil
}

func (b *BenchmarkTest01012) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest01012",
			Value:  "bar",
			MaxAge: 60 * 3,
			Secure: true,
			Path:   r.RequestURI,
			Domain: r.Host,
		})
		http.ServeFile(w, r, "sqli-02/BenchmarkTest01012.html")
		return
	}

	if r.Method == http.MethodPost {
		cookie, err := r.Cookie("BenchmarkTest01012")
		param := "noCookieValueSupplied"
		if err == nil {
			param, _ = url.QueryUnescape(cookie.Value)
		}

		bar := b.doSomething(r, param)

		sqlQuery := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

		rows, err := b.DB.Query(sqlQuery)
		if err != nil {
			http.Error(w, "Error processing request.", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var results []map[string]interface{}
		cols, _ := rows.Columns()
		for rows.Next() {
			columns := make([]interface{}, len(cols))
			for i := range columns {
				columns[i] = new(interface{})
			}
			if err := rows.Scan(columns...); err != nil {
				http.Error(w, "Error processing request.", http.StatusInternalServerError)
				return
			}
			row := make(map[string]interface{})
			for i, col := range cols {
				val := columns[i].(*interface{})
				row[col] = *val
			}
			results = append(results, row)
		}

		output, _ := json.Marshal(results)
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
}

func (b *BenchmarkTest01012) doSomething(r *http.Request, param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func main() {
	benchmark, err := NewBenchmarkTest01012()
	if err != nil {
		panic(err)
	}
	defer benchmark.DB.Close()
	http.Handle("/sqli-02/BenchmarkTest01012", benchmark)
	http.ListenAndServe(":8080", nil)
}
