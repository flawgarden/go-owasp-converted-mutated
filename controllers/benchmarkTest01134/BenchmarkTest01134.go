package controllers

import (
	"encoding/json"
	"fmt"

	"database/sql"
	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01134 struct {
	web.Controller
}

func (c *BenchmarkTest01134) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01134) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01134) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	names := c.Ctx.Request.Header
	for name := range names {
		if name == "Standard-Header" { // Replace with your condition for standard headers
			continue
		}
		param = name
		break
	}

	bar := new(Test).doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
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

type Test struct{}

func (t *Test) doSomething(param string) string {
	bar := "safe!"
	mapTest := make(map[string]interface{})
	mapTest["keyA"] = "a-Value"
	mapTest["keyB"] = param
	mapTest["keyC"] = "another-Value"
	bar = mapTest["keyB"].(string)
	return bar
}
