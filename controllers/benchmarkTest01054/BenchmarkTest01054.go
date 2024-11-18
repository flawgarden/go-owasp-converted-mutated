package controllers

import (
	"html/template"
	"net/http"
	"net/url"
)

type BenchmarkTest01054Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest01054Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ResponseWriter = w
	c.Request = r

	c.doPost()
}

func (c *BenchmarkTest01054Controller) doPost() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	if c.Request.Header.Get("Referer") != "" {
		param = c.Request.Header.Get("Referer")
	}

	param, _ = url.QueryUnescape(param)

	bar := c.doSomething(param)

	c.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.ResponseWriter.Write([]byte(bar))
}

func (c *BenchmarkTest01054Controller) doSomething(param string) string {
	bar := htmlEscape(param)
	return bar
}

func htmlEscape(input string) string {
	return template.HTMLEscapeString(input)
}
