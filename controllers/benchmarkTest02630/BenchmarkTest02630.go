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

type BenchmarkTest02630Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02630Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest02630Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest02630Controller) doPost() {
	response := c.Ctx.ResponseWriter
	request := c.Ctx.Request
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")
	queryString := request.URL.RawQuery
	paramval := "BenchmarkTest02630="
	paramLoc := -1
	if queryString != "" {
		paramLoc = indexOf(queryString, paramval)
	}
	if paramLoc == -1 {
		response.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest02630")))
		return
	}

	param := extractParam(queryString, paramLoc, paramval)
	param, _ = url.QueryUnescape(param)
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("{call %s}", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	statement, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	defer statement.Close()
	rs, err := statement.Query()
	if err != nil {
		panic(err)
	}
	defer rs.Close()
	// Process result set (implement result processing logic as needed)
}

func indexOf(query string, val string) int {
	return len(query) // Simplified for demonstration, implement actual search logic.
}

func extractParam(query string, paramLoc int, paramval string) string {
	return query[paramLoc+len(paramval):] // Simplified for demonstration.
}

func doSomething(param string) string {
	return param
}
