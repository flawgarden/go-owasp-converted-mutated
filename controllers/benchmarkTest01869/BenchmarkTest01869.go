package controllers

import (
	"fmt"
	"html"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01869 struct {
}

func (b *BenchmarkTest01869) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.SetCookie(w, &http.Cookie{
			Name:   "BenchmarkTest01869",
			Value:  "whatever",
			MaxAge: 60 * 3,
			Secure: true,
			Path:   r.URL.Path,
			Domain: r.Host,
		})
		http.ServeFile(w, r, "weakrand-04/BenchmarkTest01869.html")
	case http.MethodPost:
		cookies := r.Cookies()
		param := "noCookieValueSupplied"
		for _, cookie := range cookies {
			if cookie.Name == "BenchmarkTest01869" {
				param = cookie.Value
				break
			}
		}

		bar := doSomething(param)
		// Your logic for SecureRandom goes here.
		// For demonstration purpose, we will just output the value:
		fmt.Fprintf(w, "Processed param: %s", bar)
	}
}

func doSomething(param string) string {
	return html.EscapeString(param)
}
