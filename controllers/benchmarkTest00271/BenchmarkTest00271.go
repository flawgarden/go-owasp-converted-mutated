package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"

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

type BenchmarkTest00271Controller struct {
	beego.Controller
}

func (c *BenchmarkTest00271Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest00271Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest00271"]

	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	a75009 := param
	b75009 := fmt.Sprintf("%s SafeStuff", a75009)
	b75009 = b75009[:len(b75009)-5] + "Chars"

	map75009 := map[string]interface{}{"key75009": b75009}
	c75009 := map75009["key75009"].(string)
	d75009 := c75009[:len(c75009)-1]

	e75009 := string([]byte(d75009)) // Replace this part with actual Base64 encoding/decoding if needed
	f75009 := e75009[:len(e75009)-1]

	// Database query based on f75009 (caution: it may be vulnerable to SQL injection)
	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='%s'", f75009)
	user := models.User{}

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
