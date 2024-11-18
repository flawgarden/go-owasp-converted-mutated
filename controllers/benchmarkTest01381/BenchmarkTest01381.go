package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest01381 struct {
	beego.Controller
}

func (c *BenchmarkTest01381) Get() {
	c.doPost()
}

func (c *BenchmarkTest01381) Post() {
	c.doPost()
}

func (c *BenchmarkTest01381) doPost() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01381")
	bar := testDoSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}
	defer statement.Close()

	_, err = statement.Exec("foo")
	if err != nil {
		c.Ctx.Output.Body([]byte("Error processing request."))
		return
	}

	// Output processing logic is excluded for brevity
}

func testDoSomething(param string) string {
	bar := param
	num := 106

	if (7*42)-num > 200 {
		bar = "This should never happen"
	}

	return bar
}
