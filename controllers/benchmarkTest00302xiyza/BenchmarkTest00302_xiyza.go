//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00302/BenchmarkTest00302.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/closure.tmt with name closure_returning_closure_negative 
//Used extensions: 
//Program:
package controllers

import (
"net/url"
"os/exec"
"runtime"
"strings"
"github.com/beego/beego/v2/server/web"
"fmt"
)

type BenchmarkTest00302Controller struct {
	web.Controller
}

func (c *BenchmarkTest00302Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00302Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00302Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := c.Ctx.Request.Header["BenchmarkTest00302"]

	if len(headers) > 0 {
		param = headers[0]
		param, _ = url.QueryUnescape(param)
	}

	bar := ""
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	var cmd string
	osName := runtime.GOOS
	if strings.Contains(osName, "windows") {
		cmd = "cmd /C echo "
	} else {
		cmd = "echo "
	}

generator := makeMessageGeneratorBroken(bar)()
bar = generator("fixed_string")

	fullCmd := exec.Command(cmd + bar)
	output, err := fullCmd.CombinedOutput()
	if err != nil {
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	c.Ctx.Output.Body(output)
}

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


