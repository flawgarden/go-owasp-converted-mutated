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

type HashTestController struct {
	beego.Controller
}

func (c *HashTestController) Get() {
	c.DoPost()
}

func (c *HashTestController) Post() {
	c.DoPost()
}

func (c *HashTestController) DoPost() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00868")

	// Chain a bunch of propagators in sequence
	a := param
	b := a + " SafeStuff"
	b = b[:len(b)-len("Chars")] + "Chars"
	mapData := make(map[string]string)
	mapData["key"] = b
	c1 := mapData["key"]
	d := c1[:len(c1)-1]
	e := string(Base64Decode(Base64Encode([]byte(d))))
	f := e[:len(e)-1]

	// Connect to the database
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", f)
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

func Base64Encode(data []byte) string {
	return string(data) // Placeholder for actual encoding
}

func Base64Decode(data string) []byte {
	return []byte(data) // Placeholder for actual decoding
}
