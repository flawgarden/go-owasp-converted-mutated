package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln1Controller struct {
	web.Controller
}

func init() {
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func (c *SqlInjectionVuln1Controller) Get() {
	id := c.GetString("id")
	sqlStr := fmt.Sprintf("select * from user where id='%s'", id)
	user := models.User{}

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Database connection error")
		return
	}
	defer db.Close()

	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, "Query execution error")
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		c.CustomAbort(http.StatusInternalServerError, "JSON marshaling error")
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
