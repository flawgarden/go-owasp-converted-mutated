package controllers

import (
	"net/http"
	"strings"
)

type BenchmarkTest02131Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest02131Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r
	c.doPost()
}

func (c *BenchmarkTest02131Controller) doGet() {
	c.doPost()
}

func (c *BenchmarkTest02131Controller) doPost() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	param := c.Request.FormValue("BenchmarkTest02131")
	if param == "" {
		param = ""
	}

	bar := doSomething(c.Request, param)

	c.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	_, _ = c.ResponseWriter.Write([]byte(strings.Replace(bar, "%s", "%v", -1)))
}

func doSomething(request *http.Request, param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		bar = param[:len(param)-1]
	}

	return bar
}
