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

type BenchmarkTest00531Controller struct {
	web.Controller
}

func (c *BenchmarkTest00531Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00531Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	flag := true
	names := c.Ctx.Request.URL.Query()
	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00531" {
					param = name
					flag = false
					break
				}
			}
		}
	}

	bar := "safe!"
	map38098 := make(map[string]interface{})
	map38098["keyA-38098"] = "a-Value"
	map38098["keyB-38098"] = param
	map38098["keyC"] = "another-Value"
	bar = map38098["keyB-38098"].(string)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error fetching user"))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error marshaling user"))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
