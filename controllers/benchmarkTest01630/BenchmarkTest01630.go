package controllers

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01630Controller struct {
	web.Controller
}

func (c *BenchmarkTest01630Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01630Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	id := c.GetString("BenchmarkTest01630")
	var bar string
	if id != "" {
		bar = decodeBase64(id)
	}

	sqlStr := fmt.Sprintf("INSERT INTO users (username, password) VALUES ('foo','%s')", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec(sqlStr)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error processing request.", http.StatusInternalServerError)
		return
	}
}

func decodeBase64(param string) string {
	decoded, _ := base64.StdEncoding.DecodeString(param)
	return string(decoded)
}
