package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

type BenchmarkTest01457 struct{}

func (b *BenchmarkTest01457) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doGet(w, r)
	} else if r.Method == http.MethodPost {
		b.doPost(w, r)
	}
}

func (b *BenchmarkTest01457) doGet(w http.ResponseWriter, r *http.Request) {
	b.doPost(w, r)
}

func (b *BenchmarkTest01457) doPost(w http.ResponseWriter, r *http.Request) {
	responseWriter := http.ResponseWriter(w)
	responseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := r.URL.Query()

	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest01457" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := b.doSomething(r, param)

	http.SetCookie(w, &http.Cookie{Name: "userid", Value: bar})

	responseWriter.Write([]byte(fmt.Sprintf("Item: 'userid' with value: '%s' saved in session.", htmlEscape(bar))))
}

func (b *BenchmarkTest01457) doSomething(r *http.Request, param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}

func main() {
	http.Handle("/trustbound-00/BenchmarkTest01457", &BenchmarkTest01457{})
	http.ListenAndServe(":8080", nil)
}
