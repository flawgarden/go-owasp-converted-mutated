package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01815Controller struct {
	web.Controller
}

func (c *BenchmarkTest01815Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01815Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01815")
	bar := new(Test).doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	map72704 := make(map[string]interface{})
	map72704["keyA-72704"] = "a-Value"
	map72704["keyB-72704"] = param
	map72704["keyC"] = "another-Value"
	bar = map72704["keyB-72704"].(string)
	return bar
}
