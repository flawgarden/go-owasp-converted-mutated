package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02280Controller struct {
	web.Controller
}

func (c *BenchmarkTest02280Controller) Get() {
	param := c.GetString("BenchmarkTest02280")
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT userid from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var results int64
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	output, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe
	var bar string

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}
