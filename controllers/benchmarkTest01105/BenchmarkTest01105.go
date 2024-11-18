package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01105 struct {
	web.Controller
}

func (c *BenchmarkTest01105) Get() {
	c.Post()
}

func (c *BenchmarkTest01105) Post() {
	response := c.Ctx.ResponseWriter
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	for name := range c.Ctx.Request.Header {
		if isCommonHeader(name) {
			continue
		}
		param = name
		break
	}

	bar := new(Test).doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Write(output)
}

func isCommonHeader(header string) bool {
	commonHeaders := []string{"Content-Type", "User-Agent", "Accept"}
	for _, h := range commonHeaders {
		if strings.EqualFold(h, header) {
			return true
		}
	}
	return false
}

type Test struct{}

func (t *Test) doSomething(param string) string {
	return param
}
