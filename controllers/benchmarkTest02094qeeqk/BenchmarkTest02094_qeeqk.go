//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02094/BenchmarkTest02094.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/constructors.tmt with name class_with_string_initialization_negative 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest02094Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02094Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02094Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest02094"]
	if len(headers) > 0 {
		param = headers[0]
	}

sh := NewStringHolder()
param = sh.value

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("SQL executed successfully"))
}

func doSomething(param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}
