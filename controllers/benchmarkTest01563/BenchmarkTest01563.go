package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01563 struct {
	web.Controller
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (c *BenchmarkTest01563) Get() {
	c.Post()
}

func (c *BenchmarkTest01563) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	id := c.Ctx.Input.Query("BenchmarkTest01563")

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