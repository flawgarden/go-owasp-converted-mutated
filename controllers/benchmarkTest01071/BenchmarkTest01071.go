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

type BenchmarkTest01071 struct {
	web.Controller
}

func (c *BenchmarkTest01071) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01071) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01071) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	id := c.Ctx.Input.Header("BenchmarkTest01071")
	id = c.SanitizeInput(id)

	bar := c.processInput(id)

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

func (c *BenchmarkTest01071) processInput(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func (c *BenchmarkTest01071) SanitizeInput(param string) string {
	// Здесь можно добавить логику очистки строки
	return param
}
