package controllers

import (
	"database/sql"
	"fmt"
	"net/url"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02036 struct {
	web.Controller
}

func (c *BenchmarkTest02036) Get() {
	c.Post()
}

func (c *BenchmarkTest02036) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := ""
	headers := c.Ctx.Request.Header["BenchmarkTest02036"]
	if len(headers) > 0 {
		param = headers[0]
	}

	param, _ = url.QueryUnescape(param)

	bar := doSomething(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := fmt.Sprintf("(&(objectclass=person))(|(uid=%s)(street={0}))", bar)

	rows, err := db.Query(query, "The streetz 4 Ms bar")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err == nil {
			c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("LDAP query results:<br>Record found with name %s<br>Address: %s<br>", user.Username, user.Password)))
			found = true
		}
	}
	if !found {
		c.Ctx.ResponseWriter.Write([]byte(fmt.Sprintf("LDAP query results: nothing found for query: %s", query)))
	}
}

func doSomething(param string) string {
	bar := ""
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = append(valuesList[:0], valuesList[1:]...)
		bar = valuesList[0]
	}
	return bar
}

func main() {
	web.Router("/ldapi-00/BenchmarkTest02036", &BenchmarkTest02036{})
	web.Run()
}
