//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [22, 276, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00700/BenchmarkTest00700.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/class.tmt with name base_binary_op_negative 
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

	fileName := fmt.Sprintf("testfiles/%s", bar)
	var fos *os.File

var a12341 BaseBinaryOpClass = &DerivedBinaryOpClassDefault{}
fileName = a12341.VirtualCall(fileName, fileName)

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
