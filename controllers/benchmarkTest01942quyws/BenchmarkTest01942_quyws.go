//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01942/BenchmarkTest01942.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructorChains.tmt with name constructor_chains_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01942Controller struct {
	web.Controller
}

func (c *BenchmarkTest01942Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01942Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest01942")
	bar := doSomething(param)

cc := NewConstructorChains()
bar = cc.GetText(true)

	var cmd string
	if os.Getenv("OS") == "Windows_NT" {
		cmd = "cmd /C echo "
	} else {
		cmd = "echo "
	}

	argsEnv := "Foo=bar"
	command := exec.Command("bash", "-c", cmd+bar)
	command.Env = []string{argsEnv}

	output, err := command.CombinedOutput()
	if err != nil {
		c.Ctx.Output.Body([]byte(fmt.Sprintf("Problem executing command: %s", err.Error())))
		return
	}

	c.Ctx.Output.Body(output)
}

func doSomething(param string) string {
	return param
}
