package controllers

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02651 struct {
	web.Controller
}

func (c *BenchmarkTest02651) Get() {
	c.Post()
}

func (c *BenchmarkTest02651) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramval := "BenchmarkTest02651="
	paramLoc := -1
	if queryString != "" {
		paramLoc = getParamLoc(queryString, paramval)
	}
	if paramLoc == -1 {
		c.Ctx.Output.Body([]byte("getQueryString() couldn't find expected parameter 'BenchmarkTest02651' in query string."))
		return
	}

	param := extractParamValue(queryString, paramval, paramLoc)

	bar := doSomething(param)
	sqlStr := fmt.Sprintf("SELECT * FROM USERS where USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}

	c.Ctx.Output.Body([]byte("Query executed successfully."))
}

func getParamLoc(queryString string, paramval string) int {
	return strings.Index(queryString, paramval)
}

func extractParamValue(queryString string, paramval string, paramLoc int) string {
	param := queryString[paramLoc+len(paramval):]
	ampersandLoc := strings.Index(queryString[paramLoc:], "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramval) : paramLoc+ampersandLoc]
	}
	val, _ := url.QueryUnescape(param)
	return val
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}
	return bar
}
