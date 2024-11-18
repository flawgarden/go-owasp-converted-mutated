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

type BenchmarkTest01380 struct {
	beego.Controller
}

func (c *BenchmarkTest01380) Get() {
	c.doPost()
}

func (c *BenchmarkTest01380) Post() {
	c.doPost()
}

func (c *BenchmarkTest01380) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01380")

	bar := new(Test).doSomething(param)

	sqlStr := fmt.Sprintf("{call %s}", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
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
	num := 106
	if (7*18)+num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}
	return bar
}
