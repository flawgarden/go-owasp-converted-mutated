package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02279 struct {
	web.Controller
}

func (c *BenchmarkTest02279) Get() {
	c.Post()
}

func (c *BenchmarkTest02279) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02279")
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	var results int64
	err := db.QueryRow(sqlStr).Scan(&results)
	if err == sql.ErrNoRows {
		c.Ctx.ResponseWriter.Write([]byte("No results returned for query: " + sqlStr))
		return
	} else if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Your results are: %d", results)))
}

func doSomething(param string) string {
	bar := ""
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
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
