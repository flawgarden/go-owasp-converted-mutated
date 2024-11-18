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

type BenchmarkTest00731Controller struct {
	web.Controller
}

func (c *BenchmarkTest00731Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00731Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.GetStrings("BenchmarkTest00731")
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := param

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", bar)
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
