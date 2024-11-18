package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"xorm.io/xorm"
)

type BenchmarkTest02488Controller struct {
	xorm.Engine
}

func (c *BenchmarkTest02488Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *BenchmarkTest02488Controller) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := r.URL.Query()["BenchmarkTest02488"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(param)

	w.Header().Set("X-XSS-Protection", "0")
	obj := []interface{}{bar, "b"}
	fmt.Fprintf(w, "Formatted like: %1$s and %2$s.", obj...)
}

func doSomething(param string) string {
	bar := strings.ReplaceAll(param, "<", "&lt;")
	bar = strings.ReplaceAll(bar, ">", "&gt;")
	return bar
}
