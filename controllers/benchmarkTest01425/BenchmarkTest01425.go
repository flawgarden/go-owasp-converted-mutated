package controllers

import (
	"net/http"
)

type BenchmarkTest01425 struct{}

func (b *BenchmarkTest01425) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := r.URL.Query()

	for name, values := range names {
		if !flag {
			break
		}
		for _, value := range values {
			if value == "BenchmarkTest01425" {
				param = name
				flag = false
				break
			}
		}
	}

	bar := new(Test).doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

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
	http.Handle("/xss-02/BenchmarkTest01425", &BenchmarkTest01425{})
	http.ListenAndServe(":8080", nil)
}
