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
	sql.Open("mysql", source)
}

type BenchmarkTest00408Controller struct {
	web.Controller
}

func (c *BenchmarkTest00408Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00408Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00408")
	if param == "" {
		param = ""
	}

	bar := param

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	}

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
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
