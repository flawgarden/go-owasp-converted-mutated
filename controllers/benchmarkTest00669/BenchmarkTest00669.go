package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00669 struct {
	web.Controller
}

func (c *BenchmarkTest00669) Get() {
	c.Post()
}

func (c *BenchmarkTest00669) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	id := c.GetString("BenchmarkTest00669")
	if id == "" {
		id = ""
	}

	bar := ""
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = id
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("insert into session_data (key, value) values ('%s', '10340')", bar))
	if err != nil {
		panic(err)
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Item: '%s' with value: 10340 saved in session.", bar)))
}
