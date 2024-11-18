package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

type BenchmarkTest01689Controller struct {
	http.ResponseWriter
	*http.Request
}

func (c *BenchmarkTest01689Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01689Controller) Post() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Request.URL.RawQuery
	paramval := "BenchmarkTest01689="
	paramLoc := strings.Index(queryString, paramval)

	if paramLoc == -1 {
		c.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01689")))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(c.Request, param)

	cmd := ""
	osName := "Linux" // Replace with actual OS check if needed

	if strings.Contains(osName, "Windows") {
		cmd = "cmd /C echo "
	} else {
		cmd = "echo "
	}

	argsEnv := []string{"Foo=bar"}
	r := exec.Command(cmd+bar, argsEnv...)
	out, err := r.CombinedOutput()
	if err != nil {
		c.ResponseWriter.Write([]byte("Problem executing cmdi - TestCase"))
		return
	}

	c.ResponseWriter.Write(out)
}

type Test struct{}

func (t *Test) doSomething(r *http.Request, param string) string {
	var bar string

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	return bar
}
