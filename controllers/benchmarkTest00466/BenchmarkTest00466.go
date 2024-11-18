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

type BenchmarkTest00466Controller struct {
	web.Controller
}

func (c *BenchmarkTest00466Controller) Get() {
	c.DoPost()
}

func (c *BenchmarkTest00466Controller) Post() {
	c.DoPost()
}

func (c *BenchmarkTest00466Controller) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00466")
	bar := "safe!"
	map46344 := make(map[string]interface{})
	map46344["keyA-46344"] = "a-Value"
	map46344["keyB-46344"] = param
	map46344["keyC"] = "another-Value"
	bar = map46344["keyB-46344"].(string)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("select * from user where username='%s'", bar)
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
