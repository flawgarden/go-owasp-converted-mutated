//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01928/BenchmarkTest01928.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/map.tmt with name map_clear_negative 
//Used extensions: MACRO_Create_Map -> ~[MACRO_MapName]~ := make(map[~[TYPE@1]~]~[TYPE@2]~) | MACRO_Add_EXPR_ToMap -> ~[MACRO_MapName]~[~[EXPR_~[TYPE@1]~@1]~] = ~[EXPR_~[TYPE@2]~@2]~ | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234
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

map787234 := make(map[string]string)
map787234["NOkVp"] = "CehsI"
map787234["CehsI"] = param
map787234 = make(map[string]string)
param = map787234["CehsI"]

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
