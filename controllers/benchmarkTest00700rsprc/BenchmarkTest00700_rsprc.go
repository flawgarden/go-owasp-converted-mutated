//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [22, 276, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00700/BenchmarkTest00700.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_depth_3_array_negative 
//Used extensions: 
//Program:
package controllers

import (
	"encoding/base64"
	"fmt"
	"html"
	"os"

	"github.com/beego/beego/v2/server/web"
)

type BenchmarkTest00700Controller struct {
	web.Controller
}

func (c *BenchmarkTest00700Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00700Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")

	param := c.GetString("BenchmarkTest00700")
	bar := ""

	if param != "" {
		decoded, err := base64.StdEncoding.DecodeString(param)
		if err == nil {
			bar = string(decoded)
		}
	}

arr4124 := []string{"gwxUl"}
nested7231 := NewNestedFields3FromArray(arr4124)
bar = nested7231.nested1.nested1.nested1.values[0]

	fileName := fmt.Sprintf("testfiles/%s", bar)
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	if bar != "" {
		var err error
		fos, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Couldn't open FileOutputStream on file: '%s'\n", fileName)
			return
		}
		c.Ctx.WriteString(fmt.Sprintf("Now ready to write to file: %s", html.EscapeString(fileName)))
	}
}
