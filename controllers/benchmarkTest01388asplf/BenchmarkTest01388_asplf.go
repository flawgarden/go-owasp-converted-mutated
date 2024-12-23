package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01388Controller struct {
	web.Controller
}

func (c *BenchmarkTest01388Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01388Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01388")

value := 6
switch value {
case 1, 2, 3:
    param = "fixed_string"
case 4, 5, 6:
    param = param + "_suffix"
    fallthrough
default:
    param = "fixed_string"
}

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	var results int
	err := db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("No results returned for query: " + sqlStr))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("Your results are: " + fmt.Sprint(results)))
}

func doSomething(param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
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

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
