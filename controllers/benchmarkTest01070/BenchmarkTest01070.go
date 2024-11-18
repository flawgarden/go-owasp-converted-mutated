package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", source)
	orm.Debug = true
}

type BenchmarkTest01070 struct {
	beego.Controller
}

func (c *BenchmarkTest01070) Get() {
	c.DoPost()
}

func (c *BenchmarkTest01070) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest01070")
	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id=%s", param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

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

func main() {
	beego.Router("/weakrand-02/BenchmarkTest01070", &BenchmarkTest01070{})
	beego.Run()
}
