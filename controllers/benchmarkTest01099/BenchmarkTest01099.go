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

type BenchmarkTest01099Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01099Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest01099Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest01099Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	names := c.Ctx.Request.Header

	for name := range names {
		if !isCommonHeader(name) {
			param = name
			break
		}
	}

	bar := new(Test).doSomething(param)

	id := bar
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func isCommonHeader(name string) bool {
	commonHeaders := []string{"Accept", "User-Agent", "Content-Type"}
	for _, commonHeader := range commonHeaders {
		if name == commonHeader {
			return true
		}
	}
	return false
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	map30744 := make(map[string]interface{})
	map30744["keyA-30744"] = "a_Value"
	map30744["keyB-30744"] = param
	map30744["keyC"] = "another_Value"
	bar = map30744["keyB-30744"].(string)
	bar = map30744["keyA-30744"].(string)

	return bar
}
