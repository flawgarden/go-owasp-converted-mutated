package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

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

type BenchmarkTest01831Controller struct {
	beego.Controller
}

func (c *BenchmarkTest01831Controller) Get() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{Name: "BenchmarkTest01831", Value: "Ms+Bar", MaxAge: 60 * 3, Secure: true}
	http.SetCookie(c.Ctx.ResponseWriter, &userCookie)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "ldapi-00/BenchmarkTest01831.html")
}

func (c *BenchmarkTest01831Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	var param string
	cookies := c.Ctx.Request.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01831" {
			param = cookie.Value
			break
		}
	}
	if param == "" {
		param = "noCookieValueSupplied"
	}

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	filter := fmt.Sprintf("(&(objectclass=person))(|(uid=%s)(street={0}))", bar)
	results, err := db.Query(filter, "The streetz 4 Ms bar")
	if err != nil {
		panic(err)
	}
	defer results.Close()
	found := false
	for results.Next() {
		var user models.User
		if err := results.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			panic(err)
		}
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("LDAP query results:<br>Record found with name %s<br>", user.Username)))
		found = true
	}
	if !found {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("LDAP query results: nothing found for query: %s", filter)))
	}
}

func doSomething(param string) string {
	if param != "" {
		return string(param) // Simple example; replace with proper Base64 decoding if necessary
	}
	return ""
}
