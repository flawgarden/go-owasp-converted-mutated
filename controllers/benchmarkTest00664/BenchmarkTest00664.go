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
	web.Router("/weakrand-01/BenchmarkTest00664", &BenchmarkTest00664Controller{})
}

type BenchmarkTest00664Controller struct {
	web.Controller
}

func (c *BenchmarkTest00664Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00664Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	id := c.GetString("BenchmarkTest00664")
	if id == "" {
		id = ""
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
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
