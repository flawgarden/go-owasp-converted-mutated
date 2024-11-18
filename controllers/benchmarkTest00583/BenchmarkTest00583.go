package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"html"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00583Controller struct {
	web.Controller
}

func (c *BenchmarkTest00583Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest00583Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest00583Controller) DoPost() {
	response := c.Ctx.ResponseWriter
	request := c.Ctx.Request
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	flag := true
	names := request.URL.Query()

	for name, values := range names {
		if flag {
			for _, value := range values {
				if value == "BenchmarkTest00583" {
					param = name
					flag = false
				}
			}
		}
	}

	escapedParam := html.EscapeString(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", escapedParam)
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
