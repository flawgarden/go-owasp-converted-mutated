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

type BenchmarkTest02436Controller struct {
	web.Controller
}

func init() {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func (c *BenchmarkTest02436Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02436Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	id := c.GetString("BenchmarkTest02436")
	if id == "" {
		id = "0"
	}

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id)
	user := models.User{}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.WriteString("Error fetching user data")
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.WriteString("Error marshaling user data")
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}
