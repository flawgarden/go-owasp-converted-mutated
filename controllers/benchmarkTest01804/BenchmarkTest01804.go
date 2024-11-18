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

type BenchmarkTest01804Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01804Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01804Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest01804")
	bar := new(Test).doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME=? and PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer statement.Close()

	_, err = statement.Exec("foo")
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	// Here you can print results as required
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	map31481 := make(map[string]interface{})
	map31481["keyA-31481"] = "a_Value"
	map31481["keyB-31481"] = param
	map31481["keyC"] = "another_Value"
	bar = map31481["keyB-31481"].(string)
	bar = map31481["keyA-31481"].(string)
	return bar
}
