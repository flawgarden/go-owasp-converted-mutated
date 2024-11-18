package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00426Controller struct {
	web.Controller
}

func (c *BenchmarkTest00426Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00426Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00426")
	if param == "" {
		param = ""
	}

	var bar string
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("SET SESSION ? = ?", bar, "10340")
	if err != nil {
		panic(err)
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Item: '%s' with value: '10340' saved in session.", bar)))
}
