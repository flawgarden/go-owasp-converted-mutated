package controllers

import (
	"net/http"
)

type BenchmarkTest02690Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest02690Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r
	c.doPost()
}

func (c *BenchmarkTest02690Controller) doPost() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.getParam("BenchmarkTest02690")
	bar := doSomething(param)
	c.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.ResponseWriter.Write([]byte(bar))
}

func (c *BenchmarkTest02690Controller) getParam(param string) string {
	return c.Request.URL.Query().Get(param)
}

func doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1]

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

func main() {
	http.Handle("/xss-05/BenchmarkTest02690", &BenchmarkTest02690Controller{})
	http.ListenAndServe(":8080", nil)
}
