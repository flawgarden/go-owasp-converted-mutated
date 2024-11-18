package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00785 struct{}

func (b *BenchmarkTest00785) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	param := queryString.Get("BenchmarkTest00785")

	if param == "" {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest00785' in query string.", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", param)
	user := struct {
		Id       int
		Username string
		Password string
	}{}

	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Query execution failed", http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00785", &BenchmarkTest00785{})
	http.ListenAndServe(":8080", nil)
}
