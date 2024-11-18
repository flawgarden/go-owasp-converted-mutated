package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest01802 struct{}

func (bt *BenchmarkTest01802) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bt.doPost(w, r)
	} else if r.Method == http.MethodPost {
		bt.doPost(w, r)
	}
}

func (bt *BenchmarkTest01802) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest01802")
	bar := new(Test).doSomething(r, param)

	r.AddCookie(&http.Cookie{Name: bar, Value: "10340"})
	_, _ = fmt.Fprintf(w, "Item: '%s' with value: 10340 saved in session.", bar)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := param
	return bar
}

func main() {
	http.Handle("/trustbound-01/BenchmarkTest01802", &BenchmarkTest01802{})
	http.ListenAndServe(":8080", nil)
}
