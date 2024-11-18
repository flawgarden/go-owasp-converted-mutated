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

type BenchmarkTest00657Controller struct {
	web.Controller
}

func (c *BenchmarkTest00657Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00657Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00657")
	if param == "" {
		param = ""
	}

	bar := "safe!"
	map27260 := make(map[string]interface{})
	map27260["keyA-27260"] = "a_Value"
	map27260["keyB-27260"] = param
	map27260["keyC"] = "another_Value"
	bar = map27260["keyB-27260"].(string)
	bar = map27260["keyA-27260"].(string)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Error: %s", err)))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
