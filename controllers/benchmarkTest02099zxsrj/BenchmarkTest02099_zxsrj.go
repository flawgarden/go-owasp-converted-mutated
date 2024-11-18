//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: []
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02099/BenchmarkTest02099.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/assignment.tmt with name simple_assignment_negative 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02099Controller struct {
	web.Controller
}

func (c *BenchmarkTest02099Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02099Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest02099"]

	if len(headers) > 0 {
		param = headers[0] // just grab first element
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

bar = ""

	sqlStr := "INSERT INTO users (username, password) VALUES ('foo','" + bar + "')"

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Update complete"))
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		decoded := []byte(param)
		bar = string(decoded) // замените на вашу логику декодирования
	}
	return bar
}
