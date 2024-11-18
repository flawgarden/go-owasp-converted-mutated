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

type BenchmarkTest00354Controller struct {
	web.Controller
}

func (c *BenchmarkTest00354Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00354Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00354")
	if param == "" {
		param = ""
	}

	bar := ""
	if param != "" {
		bar = string(param)
	}

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
