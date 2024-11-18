package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00266 struct {
	web.Controller
}

func (b *BenchmarkTest00266) Get() {
	b.Post()
}

func (b *BenchmarkTest00266) Post() {
	b.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := b.Ctx.Request.Header["BenchmarkTest00266"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := "safe!"
	map42712 := make(map[string]interface{})
	map42712["keyA-42712"] = "a-Value"
	map42712["keyB-42712"] = param
	map42712["keyC"] = "another-Value"
	bar = map42712["keyB-42712"].(string)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	b.Ctx.ResponseWriter.Write(output)
}
