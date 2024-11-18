package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00352Controller struct {
	web.Controller
}

func init() {
	sql.Register("mysql", &mysql.MySQLDriver{})
}

func (c *BenchmarkTest00352Controller) Get() {
	c.HandleRequest()
}

func (c *BenchmarkTest00352Controller) Post() {
	c.HandleRequest()
}

func (c *BenchmarkTest00352Controller) HandleRequest() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00352")
	if param == "" {
		param = ""
	}

	a98424 := param
	b98424 := fmt.Sprintf("%s SafeStuff", a98424)
	c98424 := b98424[:len(b98424)-1]
	f98424 := c98424

	user := models.User{}
	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", f98424)
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}
