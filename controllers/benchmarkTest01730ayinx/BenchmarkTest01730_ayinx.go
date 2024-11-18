//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: [89]
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [89]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01730/BenchmarkTest01730.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/staticField.tmt with name class_with_static_string_field_positive 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01730Controller struct {
	web.Controller
}

func (c *BenchmarkTest01730Controller) Get() {
	c.doPost(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func (c *BenchmarkTest01730Controller) Post() {
	c.doPost(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func (c *BenchmarkTest01730Controller) doPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := request.URL.RawQuery

DEFAULT_VALUE = queryString
sfh := NewStaticFieldHolder()
queryString = sfh.value

	paramval := "BenchmarkTest01730="
	paramLoc := strings.Index(queryString, paramval)
	if paramLoc == -1 {
		response.Write([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest01730' in query string."))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	param = strings.TrimSpace(param)

	bar := NewTest().doSomething(request, param)

	sqlStr := "SELECT * from USERS where USERNAME='foo' and PASSWORD='" + bar + "'"

	db, err := sql.Open("mysql", source)
	if err != nil {
		response.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		response.Write([]byte("Error processing request."))
		return
	}
}

type Test struct{}

func NewTest() *Test {
	return &Test{}
}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := ""
	if param != "" {
		valuesList := make([]string, 0)
		valuesList = append(valuesList, "safe")
		valuesList = append(valuesList, param)
		valuesList = append(valuesList, "moresafe")

		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
