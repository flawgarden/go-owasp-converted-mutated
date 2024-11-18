package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-sec-code/models"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01993 struct {
}

func (b *BenchmarkTest01993) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := r.Header

	for name := range headers {
		if isCommonHeader(name) {
			continue
		}
		param = name
		break
	}

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(output)
}

func isCommonHeader(name string) bool {
	commonHeaders := []string{"Content-Type", "User-Agent", "Accept", "Host"} // Пример стандартных заголовков
	for _, h := range commonHeaders {
		if strings.EqualFold(name, h) {
			return true
		}
	}
	return false
}

func doSomething(param string) string {
	if param != "" {
		return param
	}
	return ""
}
