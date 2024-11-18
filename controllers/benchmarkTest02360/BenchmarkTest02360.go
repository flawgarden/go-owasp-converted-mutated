package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02360Controller struct {
	web.Controller
}

func (c *BenchmarkTest02360Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02360Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	r := c.Ctx.Request
	names := r.URL.Query()

	for name := range names {
		values := r.URL.Query()[name]
		for _, value := range values {
			if value == "BenchmarkTest02360" {
				param = name
				break
			}
		}
	}

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT userid from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	defer db.Close()

	var results int64
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		c.Ctx.WriteString("Error processing request.")
		return
	}
	c.Ctx.WriteString(fmt.Sprintf("Your results are: %d", results))
}

func doSomething(param string) string {
	bar := ""
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
