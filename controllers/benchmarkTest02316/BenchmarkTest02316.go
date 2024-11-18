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

type SqlInjectionVuln1Controller struct {
	web.Controller
}

func (c *SqlInjectionVuln1Controller) Get() {
	id := c.GetString("id")
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.Ctx.ResponseWriter.Write(output)
}
