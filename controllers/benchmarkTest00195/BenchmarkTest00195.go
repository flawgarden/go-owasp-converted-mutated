package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00195Controller struct {
	web.Controller
}

func (c *BenchmarkTest00195Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00195Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest00195")
	if param != "" {
		param, _ = url.QueryUnescape(param)
	}

	bar := param
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	user := models.User{}

	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Ctx.ResponseWriter.Write([]byte("No results can be displayed for query: " + sqlStr))
		} else {
			c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		}
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
