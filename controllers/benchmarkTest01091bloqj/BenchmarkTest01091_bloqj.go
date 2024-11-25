//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01091/BenchmarkTest01091.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/virtuality/interface.tmt with name impl_binary_op_interface_class1_negative 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01091Controller struct {
	web.Controller
}

func (c *BenchmarkTest01091Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01091Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest01091")
	param = decodeURL(param)

	bar := new(Test).doSomething(param)

var a12341 BinaryOpInterface = &ImplBinaryOpInterfaceClass1{}
bar = a12341.InterfaceCall("", bar)

	sqlStr := "SELECT * from USERS where USERNAME='foo' and PASSWORD='" + bar + "'"

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.handleError(err)
		return
	}
	defer db.Close()

	statement, err := db.Exec(sqlStr)
	if err != nil {
		c.handleError(err)
		return
	}

	c.printResults(statement)
}

func (c *BenchmarkTest01091Controller) handleError(err error) {
	c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
}

func (c *BenchmarkTest01091Controller) printResults(stmt sql.Result) {
	// Implement the logic for printing results
}

func decodeURL(param string) string {
	return strings.ReplaceAll(param, "%20", " ")
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
