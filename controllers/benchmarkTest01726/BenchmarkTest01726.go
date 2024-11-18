package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01726Controller struct {
	web.Controller
}

func (c *BenchmarkTest01726Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01726Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest01726="

	paramLoc := -1
	if queryString != "" {
		paramLoc = len(paramval)
	}
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest01726")))
		return
	}

	param := queryString[paramLoc:]
	ampersandLoc := paramLoc + len(paramval)
	if ampersandLoc < len(queryString) {
		param = queryString[paramLoc : ampersandLoc-1]
	}
	param, _ = url.QueryUnescape(param)

	bar := new(Test).doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	results, err := db.Query(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer results.Close()

	var username string
	for results.Next() {
		if err := results.Scan(&username); err == nil {
			c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("%s ", username)))
		}
	}
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	bar := ""
	guess := "ABC"
	switch guess[2] {
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
