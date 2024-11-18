package controllers

import (
	"net/http"
	"strings"

	"xorm.io/xorm"
)

type BenchmarkTest01547Controller struct {
	x *xorm.Engine
}

func (c *BenchmarkTest01547Controller) Get(w http.ResponseWriter, r *http.Request) {
	c.Post(w, r)
}

func (c *BenchmarkTest01547Controller) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01547")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(r, param)

	r.Header.Set("userid", bar)

	w.Write([]byte("Item: 'userid' with value: '" + encodeForHTML(bar) + "' saved in session."))
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	bar := ""
	if param != "" {
		bar = strings.Split(param, " ")[0]
	}
	return bar
}

func encodeForHTML(input string) string {
	// Реализация для безопасного отображения HTML
	return input
}
