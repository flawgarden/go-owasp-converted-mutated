package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00672 struct {
	web.Controller
}

func (c *BenchmarkTest00672) Get() {
	c.Post()
}

func (c *BenchmarkTest00672) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00672")
	if param == "" {
		param = ""
	}

	var bar string

	num := 86
	if (7*42)-num > 200 {
		bar = "This_should_always_happen"
	} else {
		bar = param
	}

	sqlStr := fmt.Sprintf("{call %s}", bar)

	defer func() {
		if r := recover(); r != nil {
			c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		}
	}()

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	statement, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	rs, err := statement.Query()
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	var results []models.User
	for rs.Next() {
		var user models.User
		if err := rs.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			panic(err)
		}
		results = append(results, user)
	}

	output, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}
