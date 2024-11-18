package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02699Controller struct {
	web.Controller
}

func (c *BenchmarkTest02699Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02699Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02699Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02699")
	bar := doSomething(param)

	id, err := strconv.Atoi(bar)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Invalid input", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%d", id)
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

func doSomething(param string) string {
	bar := param
	if (7*42)-106 > 200 {
		bar = "This should never happen"
	}
	return bar
}
