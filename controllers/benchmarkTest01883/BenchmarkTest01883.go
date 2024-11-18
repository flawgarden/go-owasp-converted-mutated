package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01883 struct{}

func (b *BenchmarkTest01883) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest01883",
		Value:  "bar",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.URL.Path,
		Domain: r.Host,
	})

	http.ServeFile(w, r, "sqli-04/BenchmarkTest01883.html")
}

func (b *BenchmarkTest01883) DoPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := r.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01883" {
			decodedValue, _ := url.QueryUnescape(cookie.Value)
			param = decodedValue
			break
		}
	}

	bar := doSomething(r, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	results, err := db.Query(sqlStr)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	defer results.Close()

	for results.Next() {
		var username string
		if err := results.Scan(&username); err != nil {
			http.Error(w, "Error scanning results", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, username)
	}
}

func doSomething(r *http.Request, param string) string {
	return param
}
