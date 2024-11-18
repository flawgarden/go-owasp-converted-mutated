package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02315Controller struct {
	http.Handler
}

func (c *BenchmarkTest02315Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else if r.Method == http.MethodPost {
		c.doPost(w, r)
	}
}

func (c *BenchmarkTest02315Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := r.URL.Query()
	for name, values := range names {
		if !flag {
			break
		}
		for _, value := range values {
			if value == "BenchmarkTest02315" {
				param = name
				flag = false
				break
			}
		}
	}

	bar := doSomething(r, param)

	w.Header().Set("X-XSS-Protection", "0")
	_, _ = fmt.Fprintf(w, "Formatted like: %s and %s.", "a", bar)
}

func doSomething(r *http.Request, param string) string {
	bar := param
	return bar
}
