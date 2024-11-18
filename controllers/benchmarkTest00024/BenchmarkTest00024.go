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

type SqlInjectionVuln2Controller struct {
	web.Controller
}

func (c *SqlInjectionVuln2Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00024")
	if param == "" {
		param = ""
	}

	sqlStr := fmt.Sprintf("SELECT * FROM USERS WHERE USERNAME=? AND PASSWORD='%s'", param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("foo")
	if err != nil {
		panic(err)
	}

	user := models.User{}
	err = stmt.QueryRow("foo").Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}
