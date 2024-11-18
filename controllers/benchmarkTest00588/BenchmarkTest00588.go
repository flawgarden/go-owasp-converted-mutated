package controllers

import (
	"fmt"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00588 struct{}

func (b *BenchmarkTest00588) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var param string
	flag := true
	for name, values := range r.Form {
		for _, value := range values {
			if value == "BenchmarkTest00588" {
				param = name
				flag = false
			}
		}
		if !flag {
			break
		}
	}

	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}

	session, _ := r.Cookie("session")
	if session != nil {
		http.SetCookie(w, &http.Cookie{Name: bar, Value: "10340"})
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "Item: '%s' with value: '10340' saved in session.", htmlEscape(bar))
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}

func main() {
	http.Handle("/trustbound-00/BenchmarkTest00588", &BenchmarkTest00588{})
	http.ListenAndServe(":8080", nil)
}
