//Semgrep original results: []
//Gosec original results: [78]
//CodeQL original results: [78]
//Snyk original results: []
//-------------
//Gosec analysis results: [78, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00302/BenchmarkTest00302.go
//Original file CWE's: [78]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/pointers.tmt with name pointer_init_negative 
//Used extensions: 
//Program:
package controllers

import (
	"net/url"
	"os/exec"
	"runtime"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00302Controller struct {
	web.Controller
}

func (c *BenchmarkTest00302Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00302Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00302Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := c.Ctx.Request.Header["BenchmarkTest00302"]

	if len(headers) > 0 {
		param = headers[0]
		param, _ = url.QueryUnescape(param)
	}

	bar := ""
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	var cmd string

var ptr123 *string = new(string)
var ptr234 *string = new(string)
*ptr123 = bar
ptr123 = ptr234
bar = *ptr123

	osName := runtime.GOOS
	if strings.Contains(osName, "windows") {
		cmd = "cmd /C echo "
	} else {
		cmd = "echo "
	}

	fullCmd := exec.Command(cmd + bar)
	output, err := fullCmd.CombinedOutput()
	if err != nil {
		c.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	c.Ctx.Output.Body(output)
}

func addSuffix(s *string, suf string) {
	*s = *s + suf
}

func addSuffixDoublePointer(s **string, suf *string) {
	**s = **s + *suf
}

func addSuffixDoublePointerBroken(s **string, suf *string) {
	*s = new(string)
	**s = **s + *suf
}

func swapStrings(a, b *string) {
	temp := *a
	*a = *b
	*b = temp
}

func removeSpaces(s *string) {
    *s = strings.ReplaceAll(*s, " ", "")
}


