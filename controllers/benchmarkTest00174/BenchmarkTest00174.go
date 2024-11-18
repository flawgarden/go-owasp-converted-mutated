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

type BenchmarkTest00174Controller struct {
	web.Controller
}

func (c *BenchmarkTest00174Controller) Get() {
	c.post()
}

func (c *BenchmarkTest00174Controller) Post() {
	c.post()
}

func (c *BenchmarkTest00174Controller) post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Request.Header.Get("BenchmarkTest00174")
	id := param

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error: %v", err))
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error: %v", err))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error: %v", err))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
