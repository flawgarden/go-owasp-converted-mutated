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

type BenchmarkTest00410Controller struct {
	web.Controller
}

func (c *BenchmarkTest00410Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00410Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00410")
	if param == "" {
		param = ""
	}

	bar := "safe!"
	map77851 := make(map[string]interface{})
	map77851["keyA-77851"] = "a_Value"
	map77851["keyB-77851"] = param
	map77851["keyC"] = "another_Value"
	bar = map77851["keyB-77851"].(string)
	bar = map77851["keyA-77851"].(string)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error executing query"))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error marshaling user data"))
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}
