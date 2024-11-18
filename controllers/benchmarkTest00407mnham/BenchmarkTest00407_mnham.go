//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00407/BenchmarkTest00407.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/conditional/switch.tmt with name switch_operator_any_negative 
//Used extensions: MACRO_Any_string -> ~[EXPR_string]~ | EXPR_string -> strings.TrimSpace(~[EXPR_string]~)
//Program:
package controllers

import (
	"fmt"
	"os/exec"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00407Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00407Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00407Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00407", "")
	bar := ""

	if len(param) > 0 {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	var cmd string

tmpUnique42 := ""
switch strings.TrimSpace("veUrm") {
case "YaSkg":
    bar = ""
default:
    bar = tmpUnique42
}

	var args []string

	if strings.Contains("Windows", "Windows") {
		args = []string{"cmd.exe", "/c", "echo " + bar}
	} else {
		cmd = "ls " + bar
		args = []string{"sh", "-c", cmd}
	}

	cmdEnv := []string{"foo=bar"}

	command := exec.Command(args[0], args[1:]...)
	command.Env = cmdEnv

	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing cmdi - TestCase")
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Error: %s", err.Error())))
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}
