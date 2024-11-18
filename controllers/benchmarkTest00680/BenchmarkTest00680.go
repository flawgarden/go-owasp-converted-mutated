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

type BenchmarkTest00680Controller struct {
	web.Controller
}

func (c *BenchmarkTest00680Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00680Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00680")
	if param == "" {
		param = ""
	}

	bar := "safe!"
	map33558 := make(map[string]interface{})
	map33558["keyA-33558"] = "a_Value"
	map33558["keyB-33558"] = param
	map33558["keyC"] = "another_Value"
	bar = map33558["keyB-33558"].(string)
	bar = map33558["keyA-33558"].(string)

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	if _, err := db.Exec(sqlStr); err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	user := models.User{}
	if err := db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password); err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
