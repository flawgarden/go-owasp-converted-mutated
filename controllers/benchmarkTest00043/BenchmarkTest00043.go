package controllers

import (
	"database/sql"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00043 struct {
	web.Controller
}

func (c *BenchmarkTest00043) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00043")
	if param == "" {
		param = ""
	}

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		if err != nil {
			c.Ctx.WriteString("Error processing request.")
			return
		} else {
			panic(err)
		}
	}

	c.Ctx.WriteString("Update complete")
}

func (c *BenchmarkTest00043) Get() {
	c.Post()
}
