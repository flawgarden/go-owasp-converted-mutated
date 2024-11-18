package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00025Controller struct {
	web.Controller
}

func (c *BenchmarkTest00025Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00025Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00025")
	if param == "" {
		param = ""
	}

	sqlStr := fmt.Sprintf("SELECT userid from USERS where USERNAME='foo' and PASSWORD='%s'", param)
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var results int64
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		c.Ctx.Output.Body([]byte("No results returned for query: " + sqlStr))
		return
	}
	c.Ctx.Output.Body([]byte("Your results are: " + fmt.Sprintf("%d", results)))
}
