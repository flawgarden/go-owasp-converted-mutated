package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00936 struct{}

func (b *BenchmarkTest00936) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	r = r.WithContext(ctx)

	if r.Method == http.MethodGet {
		b.doPost(w, r)
		return
	}
	if r.Method == http.MethodPost {
		b.doPost(w, r)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func (b *BenchmarkTest00936) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00936")
	var bar string

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error opening database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request.", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("No results can be displayed for query: %s<br> because the Spring batchUpdate method doesn't return results.", sqlStr)))
}

func main() {
	http.Handle("/sqli-02/BenchmarkTest00936", &BenchmarkTest00936{})
	http.ListenAndServe(":8080", nil)
}
