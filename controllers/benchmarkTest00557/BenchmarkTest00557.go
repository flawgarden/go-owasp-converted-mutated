package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"database/sql"
	"go-sec-code/models"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln1Controller struct {
	beego.Controller
}

func (c *SqlInjectionVuln1Controller) Get() {
	id := c.GetString("id")
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}