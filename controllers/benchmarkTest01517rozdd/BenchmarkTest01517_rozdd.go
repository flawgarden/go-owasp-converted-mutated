//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01517/BenchmarkTest01517.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/interface.tmt with name base_binary_op_interface_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01517Controller struct {
	web.Controller
}

func (c *BenchmarkTest01517Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01517Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01517")
	if param == "" {
		param = ""
	}

	bar := new(Test).doSomething(param)

	var a1 string
	var a2 string
	if isWindows() {
		a1 = "cmd.exe"
		a2 = "/c"
	} else {
		a1 = "sh"
		a2 = "-c"
	}

var a12341 BinaryOpInterface
if -1867566646 > 0 {
    a12341 = &ImplBinaryOpInterfaceClass1{}
} else {
    a12341 = &ImplBinaryOpInterfaceClass2{}
}
bar = a12341.InterfaceCall("", "")

	args := []string{a1, a2, "echo " + bar}

	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing cmdi - java.lang.ProcessBuilder(java.lang.String[]) Test Case")
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	c.Ctx.Output.Body(output)
}

func isWindows() bool {
	return false // заменить на реализацию проверки системы
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}
