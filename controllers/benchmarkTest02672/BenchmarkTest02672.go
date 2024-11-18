package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type HashingController struct {
	web.Controller
}

func (c *HashingController) Get() {
	c.DoPost()
}

func (c *HashingController) Post() {
	c.DoPost()
}

func (c *HashingController) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02672")
	bar := doSomething(param)

	id := bar // Vulnerable to SQL Injection
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id)
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
	map48519 := make(map[string]interface{})
	map48519["keyA-48519"] = "a_Value"
	map48519["keyB-48519"] = param
	map48519["keyC"] = "another_Value"

	bar = map48519["keyB-48519"].(string)
	bar = map48519["keyA-48519"].(string)

	return bar
}
