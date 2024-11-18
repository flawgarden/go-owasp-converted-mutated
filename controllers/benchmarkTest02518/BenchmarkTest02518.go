package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

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
	if strings.TrimSpace(id) == "" {
		c.Abort("400")
		return
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Abort("500")
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s';", id)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Abort("404")
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Abort("500")
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
