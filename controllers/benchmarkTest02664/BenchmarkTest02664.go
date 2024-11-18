package controllers

import (
	"net/http"
	"os"
)

type BenchmarkTest02664Controller struct {
	http.Handler
}

func (c *BenchmarkTest02664Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else if r.Method == http.MethodPost {
		c.doPost(w, r)
	}
}

func (c *BenchmarkTest02664Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest02664")

	bar := doSomething(param)

	fileTarget := "./" + bar + "/Test.txt"
	w.Write([]byte("Access to file: '" + fileTarget + "' created."))
	if _, err := os.Stat(fileTarget); os.IsNotExist(err) {
		w.Write([]byte(" But file doesn't exist yet."))
	} else {
		w.Write([]byte(" And file already exists."))
	}
}

func doSomething(param string) string {
	var bar string
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
