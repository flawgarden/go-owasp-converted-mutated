package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02449Controller struct {
	web.Controller
}

func (c *BenchmarkTest02449Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	c.doPost()
}

func (c *BenchmarkTest02449Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	c.doPost()
}

func (c *BenchmarkTest02449Controller) doPost() {
	param := c.GetString("BenchmarkTest02449")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("{call %s}", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer db.Close()

	result, err := db.Query(sqlStr)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer result.Close()

	// Обработка результата...
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
