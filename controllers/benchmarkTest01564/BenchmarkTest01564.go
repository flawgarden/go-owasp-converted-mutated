package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01564 struct {
	web.Controller
}

func (c *BenchmarkTest01564) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01564) Post() {
	c.DoPost()
}

func (c *BenchmarkTest01564) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01564")
	bar := escapeHTML(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func escapeHTML(input string) string {
	// Simple HTML escape implementation
	return html.EscapeString(input)
}
