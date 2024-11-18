package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest01508 struct{}

func (t *BenchmarkTest01508) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest01508")
	if param == "" {
		param = ""
	}

	bar := t.doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	w.Write([]byte(bar))
}

func (t *BenchmarkTest01508) doSomething(r *http.Request, param string) string {
	var bar string
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}

func main() {
	http.Handle("/xss-02/BenchmarkTest01508", &BenchmarkTest01508{})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
