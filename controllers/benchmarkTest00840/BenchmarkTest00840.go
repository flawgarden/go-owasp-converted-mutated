package controllers

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00840 struct {
	web.Controller
}

func (c *BenchmarkTest00840) Get() {
	c.Post()
}

func (c *BenchmarkTest00840) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Input.URI()
	paramval := "BenchmarkTest00840="
	paramLoc := -1
	if queryString != "" {
		paramLoc = len(queryString) - len(paramval) - 1
	}

	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00840")))
		return
	}

	param := queryString[paramLoc+len(paramval):]
	param = url.QueryEscape(param)

	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
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

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("No results can be displayed for query: %s<br> because the query method doesn't return results.", sqlStr)))
}
