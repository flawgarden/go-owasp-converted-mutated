package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"net/http"
	"net/url"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln2Controller struct {
	web.Controller
}

func (c *SqlInjectionVuln2Controller) Get() {
	param := c.Ctx.Input.Header("BenchmarkTest02024")
	if param == "" {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		return
	}
	param = decodeParam(param)
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}
	c.Ctx.Output.Body(output)
}

func decodeParam(param string) string {
	decoded, _ := url.QueryUnescape(param)
	return decoded
}

func doSomething(param string) string {
	if param != "" {
		return param
	}
	return ""
}
