package controllers

import (
	"net/http"
	"strings"
)

type BenchmarkTest02593Controller struct {
	http.Handler
}

func (c *BenchmarkTest02593Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.doPost(w, r)
}

func (c *BenchmarkTest02593Controller) doGet(w http.ResponseWriter, r *http.Request) {
	c.doPost(w, r)
}

func (c *BenchmarkTest02593Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := r.URL.Query().Get("BenchmarkTest02593")
	if queryString == "" {
		http.Error(w, "Couldn't find expected parameter 'BenchmarkTest02593' in query string.", http.StatusBadRequest)
		return
	}

	param := strings.TrimSpace(queryString)
	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	output := make([]interface{}, 2)
	output[0] = "a"
	output[1] = "b"
	_, _ = w.Write([]byte(bar))
}

func doSomething(param string) string {
	num := 106
	if (7*18)+num > 200 {
		return "This_should_always_happen"
	}
	return param
}
