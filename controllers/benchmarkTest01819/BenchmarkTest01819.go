package controllers

import (
	"database/sql"
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

type SqlInjectionVuln1Controller struct {
	beego.Controller
}

func (c *SqlInjectionVuln1Controller) Get() {
	id := c.GetString("BenchmarkTest01819")
	bar := new(Test).doSomething(c.Ctx.ResponseWriter, id)

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		panic(err)
	}

	c.Ctx.ResponseWriter.Write([]byte("Update complete"))
}

type Test struct{}

func (t *Test) doSomething(w http.ResponseWriter, param string) string {
	bar := "safe!"
	mapValue := make(map[string]interface{})
	mapValue["keyA-64759"] = "a-Value"
	mapValue["keyB-64759"] = param
	mapValue["keyC"] = "another-Value"
	bar = mapValue["keyB-64759"].(string)
	return bar
}
