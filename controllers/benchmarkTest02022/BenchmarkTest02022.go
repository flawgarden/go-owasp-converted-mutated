package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest02022 struct {
	beego.Controller
}

func (c *BenchmarkTest02022) Get() {
	c.Post()
}

func (c *BenchmarkTest02022) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Input.Header("BenchmarkTest02022")
	if headers != "" {
		param = headers
	}

	param, _ = url.QueryUnescape(param)
	bar := doSomething(param)

	// Simulate encryption logic
	user := models.User{}
	if err := c.getUserByID(bar, &user); err == nil {
		output, _ := json.Marshal(user)
		c.Ctx.ResponseWriter.Write(output)
	} else {
		http.Error(c.Ctx.ResponseWriter, "User not found", http.StatusNotFound)
	}
}

func (c *BenchmarkTest02022) getUserByID(id string, user *models.User) error {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id)
	return db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
}

func doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}
