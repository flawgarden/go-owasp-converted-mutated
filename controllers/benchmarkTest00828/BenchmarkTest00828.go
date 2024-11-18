package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00828 struct {
	Db *sql.DB
}

func NewBenchmarkTest00828() (*BenchmarkTest00828, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	return &BenchmarkTest00828{Db: db}, nil
}

func (bt *BenchmarkTest00828) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.RawQuery
	if queryString == "" {
		http.Error(w, "getQueryString() couldn't find expected parameter", http.StatusBadRequest)
		return
	}

	param := r.URL.Query().Get("BenchmarkTest00828")
	if param == "" {
		http.Error(w, "No parameter found", http.StatusBadRequest)
		return
	}

	if len(param) > 1 {
		param = param[:len(param)-1]
	}

	user := "SafeBystander"
	testCaseNumber := "00828"
	user += testCaseNumber
	cookieName := "rememberMe" + testCaseNumber

	cookieValue := fmt.Sprintf("%x", param)
	http.SetCookie(w, &http.Cookie{
		Name:   cookieName,
		Value:  cookieValue,
		Secure: true,
	})

	if r.Cookies() != nil {
		for _, cookie := range r.Cookies() {
			if cookieName == cookie.Name && cookie.Value == cookieValue {
				fmt.Fprintf(w, "Welcome back: %s<br/>", user)
				return
			}
		}
	}

	fmt.Fprintf(w, "%s has been remembered with cookie: %s whose value is: %s<br/>", user, cookieName, cookieValue)
}
