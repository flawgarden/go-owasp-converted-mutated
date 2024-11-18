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

type SqlInjectionVuln1Controller struct {
	beego.Controller
}

func (c *SqlInjectionVuln1Controller) Get() {
	id := c.GetString("id")
	db, err := sql.Open("mysql", source)
	if err != nil {
		c.ServeJSONError(err)
		return
	}
	defer db.Close()

	var user models.User
	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", id)
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.ServeJSONError(err)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.ServeJSONError(err)
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}

func (c *SqlInjectionVuln1Controller) ServeJSONError(err error) {
	c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
}
