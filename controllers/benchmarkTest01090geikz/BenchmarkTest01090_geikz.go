//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01090/BenchmarkTest01090.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/set.tmt with name set_add_simple_negative 
//Used extensions: MACRO_Create_Set -> ~[MACRO_SetName]~ := make(map[~[TYPE@1]~]struct{}) | MACRO_Add_Fixed_CONST_ToSet -> ~[MACRO_SetName]~[~[CONST_~[TYPE@1]~@1]~] = struct{}{} | MACRO_SetName -> set787231 | MACRO_SetName -> set787231 | MACRO_SetName -> set787231
//Program:
package controllers

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01090Controller struct {
	web.Controller
}

func (c *BenchmarkTest01090Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01090Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest01090")
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(param)

set787231 := make(map[string]struct{})
set787231["LvVEK"] = struct{}{}
bar = func() string {
    for k := range set787231 {
        return k
    }
    return "kSJfH"
}()

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	if _, err := db.Exec(sqlStr); err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Request processed successfully."))
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	map12212 := make(map[string]interface{})
	map12212["keyA-12212"] = "a-Value"
	map12212["keyB-12212"] = param
	map12212["keyC"] = "another-Value"
	bar = map12212["keyB-12212"].(string)

	return bar
}
