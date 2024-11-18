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

type BenchmarkTest00615 struct {
	web.Controller
}

func (c *BenchmarkTest00615) Get() {
	c.Post()
}

func (c *BenchmarkTest00615) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00615")
	if param == "" {
		param = ""
	}

	bar := "safe!"
	map4123 := map[string]interface{}{
		"keyA-4123": "a_Value",
		"keyB-4123": param,
		"keyC":      "another_Value",
	}
	bar = map4123["keyB-4123"].(string)
	bar = map4123["keyA-4123"].(string)

	db, err := sql.Open("mysql", source)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Database connection error"))
		return
	}
	defer db.Close()

	sqlStr := fmt.Sprintf("SELECT * FROM user WHERE id='%s'", bar)
	user := models.User{}
	err = db.QueryRow(sqlStr).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Query execution error"))
		return
	}

	output, err := json.Marshal(user)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("JSON marshal error"))
		return
	}

	c.Ctx.ResponseWriter.Write(output)
}

func main() {
	web.Router("/crypto-00/BenchmarkTest00615", &BenchmarkTest00615{})
	web.Run()
}
