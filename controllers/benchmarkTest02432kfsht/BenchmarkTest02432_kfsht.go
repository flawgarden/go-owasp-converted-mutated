//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02432/BenchmarkTest02432.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_depth_2_array_positive 
//Used extensions: 
//Program:
package controllers

import (
	"os"
	"os/exec"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02432Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02432Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02432Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02432")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

arr4124 := []string{bar}
nested7231 := NewNestedFields2FromArray(arr4124)
bar = nested7231.nested1.nested1.values[0]

	cmd := ""
	a1 := ""
	a2 := ""
	args := []string{}

	if isWindows() {
		a1 = "cmd.exe"
		a2 = "/c"
		cmd = "echo "
		args = []string{a1, a2, cmd + bar}
	} else {
		a1 = "sh"
		a2 = "-c"
		cmd = "ls "
		args = []string{a1, a2, cmd + bar}
	}

	cmdExec := exec.Command(args[0], args[1:]...)
	output, err := cmdExec.CombinedOutput()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := "safe!"
	map15850 := make(map[string]interface{})
	map15850["keyA-15850"] = "a-Value"
	map15850["keyB-15850"] = param
	map15850["keyC"] = "another-Value"
	bar = map15850["keyB-15850"].(string)
	return bar
}

func isWindows() bool {
	return os.Getenv("OS") == "Windows_NT"
}
