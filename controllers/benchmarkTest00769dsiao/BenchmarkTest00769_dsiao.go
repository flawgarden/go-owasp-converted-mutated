//Semgrep original results: [89]
//Gosec original results: [89]
//CodeQL original results: [89]
//Snyk original results: []
//-------------
//Semgrep analysis results: [89]
//Gosec analysis results: [89, 703]
//CodeQL analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00769/BenchmarkTest00769.go
//Original file CWE's: [89]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/closure.tmt with name closure_returning_closure_negative 
//Used extensions: 
//Program:
package controllers

import (
"database/sql"
"fmt"
"github.com/beego/beego/v2/server/web"
_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00769Controller struct {
	web.Controller
}

func (c *BenchmarkTest00769Controller) Get() {
	c.ServeJSON()
}

func (c *BenchmarkTest00769Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest00769")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := "safe!"
	map29572 := map[string]interface{}{
		"keyA-29572": "a-Value",
		"keyB-29572": param,
		"keyC":       "another-Value",
	}
	bar = map29572["keyB-29572"].(string)

	db, err := sql.Open("mysql", source)

generator := makeMessageGeneratorBroken(bar)()
bar = generator("fixed_string")

	if err != nil {
		c.Abort("500")
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}

	c.Ctx.WriteString("No results can be displayed for query: " + sqlStr + "<br> because the Query method doesn't return results.")
}

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


