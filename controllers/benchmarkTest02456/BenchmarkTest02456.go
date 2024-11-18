package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Database initialization code if necessary
}

type BenchmarkTest02456 struct {
	web.Controller
}

func (c *BenchmarkTest02456) Get() {
	c.Post()
}

func (c *BenchmarkTest02456) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02456")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("Update complete."))
}

func doSomething(param string) string {
	bar := ""

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
