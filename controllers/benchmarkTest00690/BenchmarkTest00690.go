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

type BenchmarkTest00690Controller struct {
	web.Controller
}

func (c *BenchmarkTest00690Controller) Get() {
	c.doPost()
}

func (c *BenchmarkTest00690Controller) Post() {
	c.doPost()
}

func (c *BenchmarkTest00690Controller) doPost() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	values := c.GetStrings("BenchmarkTest00690")
	var param string
	if len(values) > 0 {
		param = values[0]
	}

	bar := param
	switchTarget := 'C'
	switch switchTarget {
	case 'A':
		bar = param
	case 'B':
		bar = "bobs_your_uncle"
	case 'D':
		bar = param
	default:
		bar = "bobs_your_uncle"
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.Body([]byte("Database connection error"))
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.Output.Body([]byte("Query execution error"))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.Output.Body([]byte("JSON marshal error"))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
