//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00500/BenchmarkTest00500.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_depth_3_positive 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00500Controller struct {
	web.Controller
}

func (c *BenchmarkTest00500Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00500Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00500")
	var bar string

	num := 196

nested7231 := NewNestedFields3(param)
param = nested7231.nested1.nested1.nested1.value

	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	var cmd string
	if exec.Command("cmd", "/C", "echo").Run() == nil {
		cmd = "echo " + bar
	}

	argsEnv := []string{"Foo=bar"}

	r := exec.Command("bash", "-c", cmd) // Using bash for Unix-like systems
	r.Env = append(r.Env, argsEnv...)

	output, err := r.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		c.Ctx.WriteString(err.Error())
		return
	}

	c.Ctx.WriteString(string(output))
}
