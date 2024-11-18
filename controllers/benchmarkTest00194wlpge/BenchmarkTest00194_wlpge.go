//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: [563]
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00194/BenchmarkTest00194.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/varargs/varargs.tmt with name varargs_get_first_string_from_array_negative 
//Used extensions: 
//Program:
package controllers

import (
"database/sql"
"fmt"
"net/url"
beego "github.com/beego/beego/v2/server/web"
_ "github.com/go-sql-driver/mysql"
"strings"
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

bar = getFirstStringFromArray("JdjeY", "oDlHY")

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("No results can be displayed for query: " + sqlStr))
}

func getFirstString(lines ...string) string {
    return getStringWithIndex(0, lines...)
}

func getStringWithIndex(ind int, lines ...string) string {
    return lines[ind]
}

func getFirstStringFromArray(lines ...string) string {
    return lines[0]
}

func varargsWithGenerics[T any](elements ...T) T {
    return elements[0]
}

func combineStrings(strs ...string) string {
    return strings.Join(strs, ", ")
}


