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

type BenchmarkTest01994 struct {
	web.Controller
}

func (c *BenchmarkTest01994) Get() {
	c.Post()
}

func (c *BenchmarkTest01994) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	headers := c.Ctx.Request.Header
	for name := range headers {
		if isCommonHeader(name) {
			continue
		}
		param = name
		break
	}

	bar := doSomething(c.Ctx.Request, param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
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

func isCommonHeader(name string) bool {
	commonHeaders := []string{"Host", "User-Agent", "Accept", "Accept-Language", "Accept-Encoding"}
	for _, header := range commonHeaders {
		if strings.EqualFold(header, name) {
			return true
		}
	}
	return false
}

func doSomething(req *http.Request, param string) string {
	// Simulate processing of the header value
	return param
}
