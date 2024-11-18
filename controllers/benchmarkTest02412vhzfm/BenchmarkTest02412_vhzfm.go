//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02412/BenchmarkTest02412.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_depth_4_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest02412 struct {
	web.Controller
}

func (c *BenchmarkTest02412) Get() {
	c.Post()
}

func (c *BenchmarkTest02412) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Query("BenchmarkTest02412")

nested7231 := NewNestedFields4("VhVzq")
param = nested7231.nested1.nested1.nested1.nested1.value

	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	var a1, a2 string
	if strings.Contains(strings.ToLower(c.Ctx.Request.UserAgent()), "windows") {
		a1 = "cmd.exe"
		a2 = "/c"
	} else {
		a1 = "sh"
		a2 = "-c"
	}
	args := []string{a1, a2, "echo " + bar}

	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Problem executing command:", err)
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	c.Ctx.Output.Body(output)
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		bar = string(param) // Обработка параметра в соответствии с требованиями
	}

	return bar
}
