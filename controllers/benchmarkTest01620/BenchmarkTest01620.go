package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

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

type BenchmarkTest01620 struct {
	beego.Controller
}

func (c *BenchmarkTest01620) Get() {
	c.post()
}

func (c *BenchmarkTest01620) Post() {
	c.post()
}

func (c *BenchmarkTest01620) post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest01620")
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := new(Test).doSomething(param)

	sqlStr := fmt.Sprintf("{call %s}", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	result, err := db.Query(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer result.Close()

	user := models.User{}
	for result.Next() {
		if err := result.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
			return
		}
	}
	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := ""
	guess := "ABC"
	switchTarget := guess[2]

	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'C', 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	return bar
}
