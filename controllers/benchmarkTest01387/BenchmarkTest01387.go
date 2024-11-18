package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01387Controller struct {
	web.Controller
}

func (c *BenchmarkTest01387Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest01387Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest01387Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01387")

	bar := c.doSomething(param)

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	var results int
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("No results returned for query: " + sqlStr))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Your results are: " + fmt.Sprint(results)))
}

func (c *BenchmarkTest01387Controller) doSomething(param string) string {
	bar := ""

	// Simple if statement that assigns constant to bar on true condition
	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	return bar
}
