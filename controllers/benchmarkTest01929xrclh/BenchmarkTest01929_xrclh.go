//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01929/BenchmarkTest01929.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/namedreturns.tmt with name named_return_zero_values_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest01929Controller struct {
	web.Controller
}

func (c *BenchmarkTest01929Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01929Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest01929")
	bar := doSomething(param)

a123, _ := getZeroValues()
bar = a123

	var a1, a2 string
	if isWindowsOS() {
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
		fmt.Println("Problem executing command")
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := "safe!"
	map44 := make(map[string]interface{})
	map44["keyA-44"] = "a-Value"
	map44["keyB-44"] = param
	map44["keyC"] = "another-Value"
	bar = fmt.Sprintf("%v", map44["keyB-44"])
	return bar
}

func isWindowsOS() bool {
	return exec.Command("cmd", "/C", "ver").Run() == nil
}

func concat(a string, b string) (res string) {
    res = a + b
    return
}

func swap(a string, b string) (first string, second string) {
	first, second = b, a
	return
}

func brokenConcat(a string, b string) (result string) {
	defer func() {
		result = b
	}()
	result = a + b
	return
}

func getZeroValues() (x string, y string) {
    return
}


