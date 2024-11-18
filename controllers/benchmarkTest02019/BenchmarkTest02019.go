package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	// Initialize database connection (if necessary)
}

type BenchmarkTest02019 struct {
	web.Controller
}

func (c *BenchmarkTest02019) Get() {
	c.Post()
}

func (c *BenchmarkTest02019) Post() {
	response := c.Ctx.ResponseWriter
	request := c.Ctx.Request
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := request.Header["BenchmarkTest02019"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(request, param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	response.Write(output)
}

func doSomething(request *http.Request, param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}
