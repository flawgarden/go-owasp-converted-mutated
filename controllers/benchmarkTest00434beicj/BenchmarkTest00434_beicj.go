//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00434/BenchmarkTest00434.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/interface.tmt with name impl_binary_op_interface_class2_negative 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00434 struct {
	web.Controller
}

func (c *BenchmarkTest00434) Get() {
	c.Post()
}

func (c *BenchmarkTest00434) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00434")
	if param == "" {
		param = ""
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

var a12341 BinaryOpInterface = &ImplBinaryOpInterfaceClass2{}
param = a12341.InterfaceCall(param, "")

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", param)
	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("No results can be displayed for query: " + sqlStr))
}
