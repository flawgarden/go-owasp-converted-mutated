package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00435Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00435Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00435Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00435Controller) doPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00435")
	if param == "" {
		param = ""
	}

	bar := "safe!"
	map86691 := make(map[string]interface{})
	map86691["keyA-86691"] = "a-Value"
	map86691["keyB-86691"] = param
	map86691["keyC"] = "another-Value"
	bar = map86691["keyB-86691"].(string)

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='foo' AND password='%s'", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	user := models.User{}
	err = db.QueryRow("SELECT * FROM user WHERE username='foo' AND password=?", bar).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}
