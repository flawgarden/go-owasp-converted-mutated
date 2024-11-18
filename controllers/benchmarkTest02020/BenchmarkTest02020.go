package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02020 struct {
	web.Controller
}

func (c *BenchmarkTest02020) Get() {
	c.Post()
}

func (c *BenchmarkTest02020) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest02020")
	param = strings.TrimSpace(param)

	bar := doSomething(param)

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

func doSomething(param string) string {
	bar := "safe!"
	map95233 := make(map[string]interface{})
	map95233["keyA-95233"] = "a_Value"
	map95233["keyB-95233"] = param
	map95233["keyC"] = "another_Value"
	bar = map95233["keyB-95233"].(string)
	bar = map95233["keyA-95233"].(string)

	return bar
}
