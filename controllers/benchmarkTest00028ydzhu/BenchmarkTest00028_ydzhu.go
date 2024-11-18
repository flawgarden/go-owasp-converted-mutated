//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [22, 276, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00028/BenchmarkTest00028.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/list.tmt with name list_clear_start_negative 
//Used extensions: MACRO_Create_List -> ~[MACRO_ListName]~ := make([] ~[TYPE@1]~, 0) | MACRO_Add_VAR_ToList -> ~[MACRO_ListName]~ = append(~[MACRO_ListName]~, ~[VAR_~[TYPE@1]~@1]~) | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231 | MACRO_ListName -> list787231
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

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

list787231 := make([] string, 0)
list787231 = append(list787231, param)
list787231 = nil
param = list787231[0]

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
