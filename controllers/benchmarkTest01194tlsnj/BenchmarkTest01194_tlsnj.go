//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01194/BenchmarkTest01194.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_nested_string_initialization_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01194Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01194Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest01194Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest01194Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Input.Header("BenchmarkTest01194")

	if headers != "" {
		param = headers

nsh := NewNestedStringHolder("")
param = nsh.GetValue()

	}

	param, _ = url.QueryUnescape(param)

	bar := new(test).doSomething(c.Ctx.Request, param)

	var cmd string
	if strings.Contains(strings.ToLower(runtime.GOOS), "windows") {
		cmd = "cmd /C echo " + bar
	}

	argsEnv := []string{"Foo=bar"}
	r := exec.Command(cmd, argsEnv...)

	err := r.Run()
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintln("Problem executing cmdi - TestCase", err))
		return
	}
}

type test struct{}

func (t *test) doSomething(request *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}
