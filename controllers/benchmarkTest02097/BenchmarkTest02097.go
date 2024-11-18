package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02097Controller struct {
	web.Controller
}

func (c *BenchmarkTest02097Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest02097Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header.Values("BenchmarkTest02097")

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * from USERS where USERNAME='foo' and PASSWORD='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Data["json"] = "Error processing request."
		c.ServeJSON()
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		c.Data["json"] = "Error processing request."
		c.ServeJSON()
		return
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			c.Data["json"] = "Error processing request."
			c.ServeJSON()
			return
		}
		users = append(users, user)
	}

	output, err := json.Marshal(users)
	if err != nil {
		c.Data["json"] = "Error processing request."
		c.ServeJSON()
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := "safe!"
	map94015 := make(map[string]interface{})
	map94015["keyA-94015"] = "a_Value"
	map94015["keyB-94015"] = param
	map94015["keyC"] = "another_Value"
	bar = map94015["keyB-94015"].(string)
	bar = map94015["keyA-94015"].(string)

	return bar
}
