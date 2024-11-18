package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

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

type BenchmarkTest01023Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01023Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01023Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.Ctx.Request.Header.Get("BenchmarkTest01023")
	sqlStr := fmt.Sprintf("SELECT * FROM users WHERE uid='%s'", param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "User not found", http.StatusNotFound)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
