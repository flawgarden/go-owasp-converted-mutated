package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02440Controller struct {
	web.Controller
}

func (c *BenchmarkTest02440Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02440Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02440Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest02440")
	if param == "" {
		param = ""
	}

	bar := doSomething(param)

	id := bar // Здесь предполагается, что параметр используется для запроса к базе данных

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id='%s'", id)
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

func doSomething(param string) string {
	bar := "safe!"
	map85251 := make(map[string]interface{})
	map85251["keyA-85251"] = "a_Value"
	map85251["keyB-85251"] = param
	map85251["keyC"] = "another_Value"
	bar = map85251["keyB-85251"].(string)
	bar = map85251["keyA-85251"].(string)

	return bar
}
