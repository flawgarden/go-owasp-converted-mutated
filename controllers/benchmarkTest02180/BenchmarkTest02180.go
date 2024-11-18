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

type BenchmarkTest02180Controller struct {
	beego.Controller
}

func (c *BenchmarkTest02180Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02180Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02180")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	var results int64
	err := orm.NewOrm().Raw(sqlStr).QueryRow(&results)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Ctx.ResponseWriter.Write([]byte("No results returned for query: " + sqlStr))
		} else {
			c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		}
		return
	}
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Your results are: %d", results)))
}

func doSomething(param string) string {
	bar := "safe!"
	map32515 := make(map[string]interface{})
	map32515["keyA-32515"] = "a_Value"
	map32515["keyB-32515"] = param
	map32515["keyC"] = "another_Value"
	bar = map32515["keyB-32515"].(string)
	bar = map32515["keyA-32515"].(string)

	return bar
}
