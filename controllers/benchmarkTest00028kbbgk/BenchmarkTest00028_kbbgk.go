//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [22, 276, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00028/BenchmarkTest00028.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/namedreturns.tmt with name named_return_swap_simple_negative 
//Used extensions: 
//Program:
package controllers

import (
	"fmt"
	"os"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00028Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00028Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00028Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00028")

	var fileName string
	var fos *os.File

a, _ := swap(param, "suffix")
param = a

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	if param != "" {
		fileName = "/path/to/testfiles/" + param

		var err error
		fos, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Printf("Couldn't open FileOutputStream on file: '%s'\n", fileName)
			return
		}
		_, err = fos.WriteString("Now ready to write to file: " + htmlEscape(fileName))
		if err != nil {
			fmt.Printf("Error writing to file: %s\n", err)
		}
	}
}

func htmlEscape(str string) string {
	return strings.ReplaceAll(strings.ReplaceAll(str, "&", "&amp;"), "<", "&lt;")
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


