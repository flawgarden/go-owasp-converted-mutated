package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest00696Controller struct {
	http.Handler
}

func (c *BenchmarkTest00696Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest00696Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := r.FormValue("BenchmarkTest00696")
	param := values

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

	fileTarget := fmt.Sprintf("%s/%s", "testfiles", bar)
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", htmlEscape(fileTarget))))
	if _, err := os.Stat(fileTarget); err == nil {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func htmlEscape(s string) string {
	return jsonEscape(s) // or use a proper HTML escape function
}

func jsonEscape(s string) string {
	return fmt.Sprintf("%q", s)[1 : len(fmt.Sprintf("%q", s))-1] // primitive implementation
}
