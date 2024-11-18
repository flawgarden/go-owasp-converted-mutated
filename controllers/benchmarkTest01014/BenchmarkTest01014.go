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

type BenchmarkTest struct{}

func (bt *BenchmarkTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest01014",
			Value:  "2222",
			MaxAge: 60 * 3,
			Secure: true,
			Path:   r.URL.Path,
			Domain: r.URL.Host,
		})
		http.ServeFile(w, r, "./xpathi-00/BenchmarkTest01014.html")
	}
	if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest01014" {
				param, _ = url.QueryUnescape(cookie.Value)
				break
			}
		}
		bar := bt.doSomething(r, param)

		db, err := sql.Open("mysql", source)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var result string
		sqlStr := fmt.Sprintf("SELECT value FROM employees WHERE id='%s'", bar)
		err = db.QueryRow(sqlStr).Scan(&result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := map[string]string{"result": result}
		output, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
}

func (bt *BenchmarkTest) doSomething(r *http.Request, param string) string {
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	var bar string
	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}
	return bar
}

func main() {
	http.Handle("/xpathi-00/BenchmarkTest01014", &BenchmarkTest{})
	http.ListenAndServe(":8080", nil)
}
