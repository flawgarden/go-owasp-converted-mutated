//Semgrep original results: []
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Gosec analysis results: [89, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01559/BenchmarkTest01559.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/conditional/for.tmt with name for_operator_strcpy_positive 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01559Controller struct {
	web.Controller
}

func (c *BenchmarkTest01559Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01559Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01559")
	if param == "" {
		param = ""
	}

tmpUnique42 := param
param = ""
for i := 0; i < len(tmpUnique42); i++ {
    c := tmpUnique42[i]
    param += string(c)
}

	bar := new(Test).DoSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer statement.Close()

	results, err := statement.Exec()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	rowsAffected, _ := results.RowsAffected()
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Rows affected: %d", rowsAffected)))
}

type Test struct{}

func (t *Test) DoSomething(request *http.Request, param string) string {
	bar := "safe!"
	map31651 := make(map[string]interface{})
	map31651["keyA-31651"] = "a-Value"
	map31651["keyB-31651"] = param
	map31651["keyC"] = "another-Value"
	bar = map31651["keyB-31651"].(string)

	return bar
}
