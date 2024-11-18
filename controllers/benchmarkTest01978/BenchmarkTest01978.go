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

type BenchmarkTest01978 struct {
	web.Controller
}

func (c *BenchmarkTest01978) Get() {
	c.Post()
}

func (c *BenchmarkTest01978) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	for _, name := range c.Ctx.Request.Header.Values("HeaderName") {
		if commonHeadersContains(name) {
			continue
		}
		param = name
		break
	}

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", bar)
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
	if param != "" && len(param) > 1 {
		sbxyz6576 := []rune(param)
		bar = string(append(sbxyz6576[:len(param)-1], 'Z'))
	}
	return bar
}

func commonHeadersContains(header string) bool {
	commonHeaders := []string{"Content-Type", "Authorization"}
	for _, commonHeader := range commonHeaders {
		if header == commonHeader {
			return true
		}
	}
	return false
}
