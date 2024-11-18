package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01385Controller struct {
	web.Controller
}

func (c *BenchmarkTest01385Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01385Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01385")
	bar := new(Test).doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	connection, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer connection.Close()

	statement, err := connection.Prepare(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, err = statement.Exec("foo")
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}

	// Print results logic
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[1]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = param
	default:
		bar = "bob's your uncle"
	}

	return bar
}
