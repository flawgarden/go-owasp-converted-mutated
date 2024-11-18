package controllers

import (
	"fmt"
	"net/http"
)

type BenchmarkTest02302Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest02302Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	flag := true
	for name, values := range r.Form {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest02302" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := doSomething(r, param)

	fileTarget := fmt.Sprintf("%s/Test.txt", bar)
	w.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))
	if fileExists(fileTarget) {
		w.Write([]byte(" And file already exists."))
	} else {
		w.Write([]byte(" But file doesn't exist yet."))
	}
}

func doSomething(r *http.Request, param string) string {
	// Mocking the behavior of the original Java method
	return fmt.Sprintf("/path/to/resource/%s", param)
}

func fileExists(filename string) bool {
	// Mock implementation for the sake of demonstrating the example
	// In a real application, you should check if the file exists.
	return false
}
