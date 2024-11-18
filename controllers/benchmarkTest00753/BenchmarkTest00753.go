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

type BenchmarkTest00753Controller struct {
	web.Controller
}

func init() {
	// MySQL Driver registration if needed
}

func (c *BenchmarkTest00753Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00753Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	values := c.Ctx.Input.Query("BenchmarkTest00753")
	param := ""
	if values != "" {
		param = values
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := param
	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id)
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

	rememberMeKey := "someRandomValue" // replace with actual secure random value generation logic
	c.SetSession("rememberMe", rememberMeKey)
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("User remembered with key: %s<br/>", rememberMeKey)))
}
