//Semgrep original results: [89]
//Gosec original results: []
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: [89]
//CodeQL analysis results: []
//Gosec analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest02449/BenchmarkTest02449.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/map.tmt with name map_clear_negative 
//Used extensions: MACRO_Create_Map -> ~[MACRO_MapName]~ := make(map[~[TYPE@1]~]~[TYPE@2]~) | MACRO_Add_EXPR_ToMap -> ~[MACRO_MapName]~[~[EXPR_~[TYPE@1]~@1]~] = ~[EXPR_~[TYPE@2]~@2]~ | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234
//Program:
package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02449Controller struct {
	web.Controller
}

func (c *BenchmarkTest02449Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	c.doPost()
}

func (c *BenchmarkTest02449Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	c.doPost()
}

func (c *BenchmarkTest02449Controller) doPost() {
	param := c.GetString("BenchmarkTest02449")
	if param == "" {
		param = ""
	}

map787234 := make(map[string]string)
map787234["DrNnl"] = param
map787234[param] = param
map787234 = make(map[string]string)
param = map787234[param]

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("{call %s}", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer db.Close()

	result, err := db.Query(sqlStr)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer result.Close()

	// Обработка результата...
}

func doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}
