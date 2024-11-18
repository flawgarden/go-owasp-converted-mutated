package controllers

import (
	"database/sql"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01853 struct{}

func (b *BenchmarkTest01853) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.SetCookie(w, &http.Cookie{
			Name:     "BenchmarkTest01853",
			Value:    "whatever",
			MaxAge:   60 * 3,
			Secure:   true,
			Path:     r.URL.Path,
			HttpOnly: false,
			Domain:   r.URL.Hostname(),
		})
		http.ServeFile(w, r, "weakrand-04/BenchmarkTest01853.html")
		return
	}

	if r.Method == http.MethodPost {
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest01853" {
				param, _ = url.QueryUnescape(cookie.Value)
				break
			}
		}

		bar := doSomething(param)

		db, err := sql.Open("mysql", source)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		_ = db.Ping()

		w.Write([]byte("Processed: " + bar + "<br/>"))
	}
}

func doSomething(param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func main() {
	http.Handle("/weakrand-04/BenchmarkTest01853", &BenchmarkTest01853{})
	http.ListenAndServe(":8080", nil)
}
