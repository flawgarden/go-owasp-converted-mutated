package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02517Controller struct {
	web.Controller
}

func (c *BenchmarkTest02517Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest02517Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest02517Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest02517")

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, fmt.Sprintf("Query error: %v", err), http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, fmt.Sprintf("JSON error: %v", err), http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	// Simulating the behavior of the original doSomething method
	return param
}
