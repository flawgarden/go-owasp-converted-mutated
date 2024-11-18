package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-sec-code/models"
	"log"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest00445 struct {
	web.Controller
}

func (c *BenchmarkTest00445) Get() {
	c.ServeJSON()
}

func (c *BenchmarkTest00445) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	id := c.GetString("BenchmarkTest00445")
	if id == "" {
		log.Println("No input received")
		return
	}

	sqlStr := fmt.Sprintf("select * from user where id=%s", id)
	user, err := queryUser(sqlStr)
	if err != nil {
		log.Println("Database query error:", err)
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		log.Println("JSON marshal error:", err)
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}

func queryUser(query string) (models.User, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	user := models.User{}
	err = db.QueryRow(query).Scan(&user.Id, &user.Username, &user.Password)
	return user, err
}

func main() {
	web.Router("/crypto-00/BenchmarkTest00445", &BenchmarkTest00445{})
	web.Run()
}
