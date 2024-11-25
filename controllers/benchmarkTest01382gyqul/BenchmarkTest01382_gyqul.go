//Original file region: null, null, null, null
//Mutated file region: null, null, null, null
//Semgrep original results: []
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [563]
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01382/BenchmarkTest01382.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/conditional/dowhile.tmt with name do_while_operator_negative 
//Used extensions: EXPR_bool -> strings.EqualFold(~[EXPR_string]~, ~[EXPR_string]~) | EXPR_string -> strings.Replace(~[EXPR_string]~, "/", ".", -1) | EXPR_string -> strings.Join([]string{~[EXPR_string]~, ~[EXPR_string]~}, "")
//Program:
package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01382Controller struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

func (c *BenchmarkTest01382Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01382Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01382")
	bar := new(Test).doSomething(c.Ctx.Request, param)

for {
    bar = ""
    if !strings.EqualFold(strings.Replace("oPfbw", "/", ".", -1), strings.Join([]string{param, "CTyse"}, "")) {
        break
    }
}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, err = statement.Exec("foo")
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return strings.TrimSpace(bar)
}
