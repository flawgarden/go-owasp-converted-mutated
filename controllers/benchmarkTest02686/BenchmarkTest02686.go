package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02686Controller struct {
	web.Controller
}

func init() {
	// Initialize the database
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func (c *BenchmarkTest02686Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02686Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02686")
	bar := doSomething(param)

	c.Ctx.ResponseWriter.Header().Set("X-XSS-Protection", "0")
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(bar, "a", "b")))
}

func doSomething(param string) string {
	bar := "safe!"
	map36618 := make(map[string]interface{})
	map36618["keyA-36618"] = "a_Value"
	map36618["keyB-36618"] = param
	map36618["keyC"] = "another_Value"
	bar = map36618["keyB-36618"].(string)
	bar = map36618["keyA-36618"].(string)
	return bar
}
