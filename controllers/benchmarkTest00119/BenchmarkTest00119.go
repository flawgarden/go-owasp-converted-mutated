package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

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

type BenchmarkTest struct {
	beego.Controller
}

func (c *BenchmarkTest) Get() {
	c.DoRequestHandler()
}

func (c *BenchmarkTest) Post() {
	c.DoRequestHandler()
}

func (c *BenchmarkTest) DoRequestHandler() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.Ctx.Input.Header("BenchmarkTest00119")
	param, _ = url.QueryUnescape(param)

	a := param
	b := strings.Builder{}
	b.WriteString(a)
	b.WriteString(" SafeStuff")
	b.WriteString("Chars")

	mapData := map[string]interface{}{"key": b.String()}
	cString := mapData["key"].(string)
	d := cString[:len(cString)-1]
	f := d

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE username='%s'", f)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
