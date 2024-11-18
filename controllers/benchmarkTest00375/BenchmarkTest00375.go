package controllers

import (
	"fmt"
	"net/http"
	"strings"
)

type BenchmarkTest00375Controller struct {
	http.Handler
}

func (c *BenchmarkTest00375Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest00375")
	if param == "" {
		param = ""
	}

	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{"a", "b"}
	output := formatOutput(bar, obj)
	w.Write([]byte(output))
}

func formatOutput(bar string, obj []interface{}) string {
	return fmt.Sprintf("%s %v", bar, obj)
}
