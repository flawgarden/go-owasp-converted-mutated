package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	sql.Register("mysql", &mysql.MySQLDriver{})
}

type BenchmarkTest00309 struct {
	web.Controller
}

func (c *BenchmarkTest00309) Get() {
	c.Post()
}

func (c *BenchmarkTest00309) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest00309"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := "safe!"
	map92785 := map[string]interface{}{
		"keyA-92785": "a_Value",
		"keyB-92785": param,
		"keyC":       "another_Value",
	}
	bar = map92785["keyB-92785"].(string)
	bar = map92785["keyA-92785"].(string)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error opening database"))
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("User not found"))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error marshalling JSON"))
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
