package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-sec-code/models"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00075Controller struct {
	web.Controller
}

func (c *BenchmarkTest00075Controller) Get() {
	http.SetCookie(c.Ctx.ResponseWriter, &http.Cookie{
		Name:   "BenchmarkTest00075",
		Value:  "someSecret",
		Path:   c.Ctx.Request.RequestURI,
		MaxAge: 60 * 3,
		Secure: true,
	})

	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "hash-00/BenchmarkTest00075.html")
}

func (c *BenchmarkTest00075Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := c.Ctx.Request.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00075" {
			param = cookie.Value
			break
		}
	}

	bar := param
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	}

	id := strings.TrimSpace(bar)
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", id)
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
