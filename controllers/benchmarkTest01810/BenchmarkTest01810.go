package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01810Controller struct {
	web.Controller
}

func init() {
	web.Router("/sqli-03/BenchmarkTest01810", &BenchmarkTest01810Controller{})
}

func (c *BenchmarkTest01810Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01810Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01810")
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT TOP 1 userid from USERS where USERNAME='foo' and PASSWORD='%s'", bar)

	results, err := queryForMap(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	output, err := json.Marshal(results)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := param
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}
	return bar
}

func queryForMap(sqlStr string) (map[string]interface{}, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow(sqlStr)

	var results map[string]interface{}
	if err := row.Scan(&results); err != nil {
		return nil, err
	}
	return results, nil
}
