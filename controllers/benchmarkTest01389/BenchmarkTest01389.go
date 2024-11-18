package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01389 struct {
	web.Controller
}

func (c *BenchmarkTest01389) Get() {
	c.doPost()
}

func (c *BenchmarkTest01389) Post() {
	c.doPost()
}

func (c *BenchmarkTest01389) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01389")

	bar := new(Test).doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	list, err := queryDatabase(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Your results are: <br>"))
	for _, o := range list {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("%s<br>", o)))
	}
}

func queryDatabase(sqlStr string) ([]map[string]interface{}, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	columns, _ := rows.Columns()
	for rows.Next() {
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(interface{})
		}
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			rowMap[col] = *(values[i].(*interface{}))
		}
		results = append(results, rowMap)
	}
	return results, nil
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
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
