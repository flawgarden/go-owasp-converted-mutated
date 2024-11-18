package controllers

import (
	"fmt"

	"database/sql"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02446Controller struct {
	web.Controller
}

func (c *BenchmarkTest02446Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02446Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.Ctx.Input.Query("BenchmarkTest02446")

	if param == "" {
		param = "0"
	}

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("INSERT INTO session_data (key, value) VALUES ('%s', '%d')", bar, 10340)
	_, err = db.Exec(sqlStr)
	if err != nil {
		panic(err)
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Item: '%s' with value: 10340 saved in session.", bar)))
}

func doSomething(param string) string {
	bar := ""
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func main() {
	web.Router("/trustbound-01/BenchmarkTest02446", &BenchmarkTest02446Controller{})
	web.Run()
}
