package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00798 struct {
}

func (b *BenchmarkTest00798) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	paramval := "BenchmarkTest00798="
	paramLoc := -1

	if queryString != "" {
		paramLoc = findParamLocation(queryString, paramval)
	}

	if paramLoc == -1 {
		http.Error(w, "getQueryString() couldn't find expected parameter 'BenchmarkTest00798' in query string.", http.StatusBadRequest)
		return
	}

	param := extractParam(queryString, paramLoc, paramval)
	bar := param

	if param != "" && len(param) > 1 {
		bar = replaceLastChar(param, 'Z')
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshalling user data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func findParamLocation(queryString, paramval string) int {
	return -1 // Logic to find parameter location
}

func extractParam(queryString string, paramLoc int, paramval string) string {
	return "" // Logic to extract parameter value
}

func replaceLastChar(s string, newChar rune) string {
	return s[:len(s)-1] + string(newChar)
}
