package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01556 struct {
	web.Controller
}

func (c *BenchmarkTest01556) Get() {
	c.Post()
}

func (c *BenchmarkTest01556) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01556")
	if param == "" {
		param = ""
	}
	bar := new(Test).doSomething(c.Ctx.Request, param)
	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	results, err := queryDatabase(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Your results are: %d", results)))
}

func queryDatabase(query string) (int64, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var results int64
	err = db.QueryRow(query).Scan(&results)
	if err != nil {
		return 0, err
	}
	return results, nil
}

type Test struct{}

func (t *Test) doSomething(request *http.Request, param string) string {
	var bar string
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

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
