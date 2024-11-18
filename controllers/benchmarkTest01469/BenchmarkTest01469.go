package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

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

type BenchmarkTest01469 struct {
	beego.Controller
}

func (c *BenchmarkTest01469) Get() {
	c.Post()
}

func (c *BenchmarkTest01469) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Request.URL.Query()
	for name := range names {
		if flag {
			values := names[name]
			for i := 0; i < len(values) && flag; i++ {
				value := values[i]
				if value == "BenchmarkTest01469" {
					param = name
					flag = false
				}
			}
		}
	}

	bar := new(Test).doSomething(c.Ctx.Request, param)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	list, err := queryForList(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("Your results are: <br>"))
	for _, item := range list {
		output, _ := json.Marshal(item)
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("%s<br>", output)))
	}
}

func queryForList(sqlStr string) ([]map[string]interface{}, error) {
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

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	results := make([]map[string]interface{}, 0)
	for rows.Next() {
		row := make([]interface{}, len(columns))
		rowPointers := make([]interface{}, len(columns))
		for i := range row {
			rowPointers[i] = &row[i]
		}

		if err := rows.Scan(rowPointers...); err != nil {
			return nil, err
		}

		result := make(map[string]interface{})
		for i, col := range columns {
			result[col] = row[i]
		}
		results = append(results, result)
	}

	return results, nil
}

type Test struct{}

func (t *Test) doSomething(req *http.Request, param string) string {
	a83916 := param
	b83916 := a83916 + " SafeStuff"
	return b83916[:len(b83916)-1]
}
