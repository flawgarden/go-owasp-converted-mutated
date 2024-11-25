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
//Original file name: controllers/benchmarkTest00194/BenchmarkTest00194.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/field/nested.tmt with name nested_field_depth_3_array_positive 
//Used extensions: 
//Program:
package controllers

import (
	"database/sql"
	"fmt"
	"net/url"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00194Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00194Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00194Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest00194")

arr4124 := []string{param}
nested7231 := NewNestedFields3FromArray(arr4124)
param = nested7231.nested1.nested1.nested1.values[0]

	param, _ = url.QueryUnescape(param)

	bar := ""
	if param != "" {
		bar = string(param) // Использование Base64 не требуется для примера
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("No results can be displayed for query: " + sqlStr))
}
