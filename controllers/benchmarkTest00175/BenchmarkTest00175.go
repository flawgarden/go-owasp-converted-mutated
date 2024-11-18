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

func init() {
	// Initialize database connection
}

type BenchmarkTest00175Controller struct {
	web.Controller
}

func (c *BenchmarkTest00175Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00175Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.Ctx.Request.Header.Get("BenchmarkTest00175")

	bar := "safe!"
	map50591 := make(map[string]interface{})
	map50591["keyA-50591"] = "a_Value"
	map50591["keyB-50591"] = param
	map50591["keyC"] = "another_Value"
	bar = map50591["keyB-50591"].(string)
	bar = map50591["keyA-50591"].(string)

	sqlStr := fmt.Sprintf("select * from user where id=%s", bar)
	user := models.User{}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
