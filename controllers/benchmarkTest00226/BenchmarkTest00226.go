package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00226Controller struct {
	web.Controller
}

func (c *BenchmarkTest00226Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00226Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	for name, values := range c.Ctx.Request.Header {
		if !isCommonHeader(name) {
			param = strings.Join(values, ", ")
			break
		}
	}

	bar := "safe!"
	map56895 := map[string]interface{}{
		"keyA-56895": "a_Value",
		"keyB-56895": param,
		"keyC":       "another_Value",
	}
	bar = map56895["keyB-56895"].(string)
	bar = map56895["keyA-56895"].(string)

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
	commonHeaders := []string{"User-Agent", "Accept", "Connection"} // Add more common headers as needed
	for _, h := range commonHeaders {
		if strings.EqualFold(h, name) {
			return true
		}
	}
	return false
}
