package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00755Controller struct {
	web.Controller
}

func (c *BenchmarkTest00755Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00755Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest00755")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	var bar string
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO session (userid) VALUES (?)", bar)
	if err != nil {
		panic(err)
	}

	output := fmt.Sprintf("Item: 'userid' with value: '%s' saved in session.", bar)
	c.Ctx.ResponseWriter.Write([]byte(output))
}
