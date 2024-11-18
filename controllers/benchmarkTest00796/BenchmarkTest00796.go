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

type BenchmarkTest00796 struct {
	web.Controller
}

func (c *BenchmarkTest00796) Get() {
	c.Post()
}

func (c *BenchmarkTest00796) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	queryString := c.Ctx.Request.URL.RawQuery
	paramVal := "BenchmarkTest00796="
	paramLoc := -1
	if queryString != "" {
		paramLoc = indexOf(queryString, paramVal)
	}
	if paramLoc == -1 {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("getQueryString() couldn't find expected parameter '%s' in query string.", "BenchmarkTest00796")))
		return
	}

	param := queryString[paramLoc+len(paramVal):]
	ampersandLoc := indexOf(queryString, "&")
	if ampersandLoc != -1 {
		param = queryString[paramLoc+len(paramVal) : ampersandLoc]
	}

	bar := "safe!"
	map72213 := make(map[string]interface{})
	map72213["keyA-72213"] = "a-Value"
	map72213["keyB-72213"] = param
	map72213["keyC"] = "another-Value"
	bar = map72213["keyB-72213"].(string)

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
	c.Ctx.ResponseWriter.Write(output)
}

func indexOf(s, substr string) int {
	if len(substr) > len(s) {
		return -1
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
