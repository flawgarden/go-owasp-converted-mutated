package controllers

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00196Controller struct {
	web.Controller
}

func (c *BenchmarkTest00196Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00196Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest00196")
	bar := ""

	if param != "" {
		param, _ = url.QueryUnescape(param)
	}

	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("No results can be displayed for query: " + sqlStr))
}
