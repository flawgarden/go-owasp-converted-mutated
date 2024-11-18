package controllers

import (
	"database/sql"
	"fmt"

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

type BenchmarkTest00619Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00619Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00619Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00619")
	if param == "" {
		param = ""
	}

	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	fileTarget := fmt.Sprintf("%s/%s", "testfiles_dir", bar)
	c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("Access to file: '%s' created.", fileTarget)))
	if _, err := sql.Open("mysql", source); err == nil {
		c.Ctx.ResponseWriter.Write([]byte(" And file already exists."))
	} else {
		c.Ctx.ResponseWriter.Write([]byte(" But file doesn't exist yet."))
	}
}
