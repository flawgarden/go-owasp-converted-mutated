//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01928/BenchmarkTest01928.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/closure.tmt with name closure_capturing_negative 
//Used extensions: 
//Program:
package controllers

import (
"fmt"
"os/exec"
"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01928Controller struct {
	web.Controller
}

func (c *BenchmarkTest01928Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01928Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01928Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest01928")

addPrefix := makePrefixer("tmRyT")
tmp123 := addPrefix("_suffix")
param = tmp123

	bar := doSomething(param)

	var argList []string
	if isWindows() {
		argList = append(argList, "cmd.exe", "/c")
	} else {
		argList = append(argList, "sh", "-c")
	}
	argList = append(argList, "echo "+bar)

	cmd := exec.Command(argList[0], argList[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing command:", err)
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
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

func isWindows() bool {
	return exec.Command("cmd.exe").Run() == nil
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


