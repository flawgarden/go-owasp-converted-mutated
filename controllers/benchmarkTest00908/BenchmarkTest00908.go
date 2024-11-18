package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

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

type BenchmarkTest00908Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00908Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00908Controller) Post() {
	c.Ctx.Output.Header("Content-Type", "text/html;charset=UTF-8")
	id := c.GetString("BenchmarkTest00908")

	// Safe guess
	guess := "ABC"
	switchTarget := guess[1] // condition 'B', which is safe

	var bar string
	switch switchTarget {
	case 'A':
		bar = id
	case 'B':
		bar = "bob"
	case 'C', 'D':
		bar = id
	default:
		bar = "bob's your uncle"
	}

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where id=%s", bar)
	user := models.User{}
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
