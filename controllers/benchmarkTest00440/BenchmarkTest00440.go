package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00440Controller struct {
	web.Controller
}

func (c *BenchmarkTest00440Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00440Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00440")
	if param == "" {
		param = ""
	}

	bar := "safe!"
	map67409 := make(map[string]interface{})
	map67409["keyA-67409"] = "a_Value"
	map67409["keyB-67409"] = param
	map67409["keyC"] = "another_Value"
	bar = map67409["keyB-67409"].(string)
	bar = map67409["keyA-67409"].(string)

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("Update complete."))
}
