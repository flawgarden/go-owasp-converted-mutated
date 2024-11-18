//Semgrep original results: [79]
//Gosec original results: []
//CodeQL original results: [79]
//Snyk original results: [79]
//-------------
//Semgrep analysis results: [79]
//CodeQL analysis results: []
//Snyk analysis results: []
//Gosec analysis results: []
//Original file name: controllers/benchmarkTest02583/BenchmarkTest02583.go
//Original file CWE's: [79]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/concurrency/channels.tmt with name channel_struct_negative 
//Used extensions: 
//Program:
package controllers

import (
"net/http"
"net/url"
beego "github.com/beego/beego/v2/server/web"
"sync"
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

dataChannel := make(chan Data, 1)
dataChannel <- Data{Value: queryString}

var wg sync.WaitGroup
wg.Add(1)

go func() {
    data := <-dataChannel
    data.Value = "constant_string"
    dataChannel <- data
}()

wg.Wait()

readData := <-dataChannel
queryString = readData.Value

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
