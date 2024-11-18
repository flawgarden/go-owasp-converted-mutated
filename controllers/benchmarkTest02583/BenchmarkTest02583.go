package controllers

import (
	"net/http"
	"net/url"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02583 struct {
	beego.Controller
}

func (c *BenchmarkTest02583) Get() {
	c.post(c.Ctx.Request, c.Ctx.ResponseWriter)
}

func (c *BenchmarkTest02583) Post() {
	c.post(c.Ctx.Request, c.Ctx.ResponseWriter)
}

func (c *BenchmarkTest02583) post(req *http.Request, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := req.URL.RawQuery
	paramval := "BenchmarkTest02583="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval) - 1
	}
	if paramLoc == -1 {
		http.Error(res, "getQueryString() couldn't find expected parameter 'BenchmarkTest02583' in query string.", http.StatusBadRequest)
		return
	}

	param, err := url.QueryUnescape(queryString[paramLoc+len(paramval):])
	if err != nil {
		http.Error(res, "Invalid parameter", http.StatusBadRequest)
		return
	}

	bar := doSomething(req, param)
	res.Header().Set("X-XSS-Protection", "0")
	res.Write([]byte(bar))
}

func doSomething(req *http.Request, param string) string {
	bar := param
	return bar
}
