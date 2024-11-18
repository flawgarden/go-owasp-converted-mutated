package controllers

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01086 struct {
	web.Controller
}

func (c *BenchmarkTest01086) Get() {
	c.Post()
}

func (c *BenchmarkTest01086) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest01086")
	if param == "" {
		param = ""
	}

	param, _ = url.QueryUnescape(param)

	bar := c.doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	if err := executeSQL(sqlStr); err != nil {
		c.Ctx.Output.Body([]byte("Error processing request."))
	} else {
		c.Ctx.Output.Body([]byte("No results can be displayed for query: " + escapeHTML(sqlStr)))
	}
}

func (c *BenchmarkTest01086) doSomething(param string) string {
	bar := ""

	num := 106

	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}

func executeSQL(sqlStr string) error {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	return err
}

func escapeHTML(str string) string {
	return fmt.Sprintf("%s", str) // Replace with actual escaping logic if needed
}
