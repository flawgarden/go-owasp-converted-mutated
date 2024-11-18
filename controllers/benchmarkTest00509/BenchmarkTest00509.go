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

type BenchmarkTest00509 struct {
	web.Controller
}

func (c *BenchmarkTest00509) Get() {
	c.Post()
}

func (c *BenchmarkTest00509) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00509", "")
	bar := "safe!"
	map94841 := make(map[string]interface{})
	map94841["keyA-94841"] = "a_Value"
	map94841["keyB-94841"] = param
	map94841["keyC"] = "another_Value"
	bar = map94841["keyB-94841"].(string)
	bar = map94841["keyA-94841"].(string)

	sqlStr := fmt.Sprintf("{call %s}", bar)

	db, err := sql.Open("mysql", source)
	if err != nil {
		log.Println("Error opening database:", err)
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer db.Close()

	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Println("Error executing query:", err)
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Password); err != nil {
			log.Println("Error scanning row:", err)
			break
		}
	}

	output, err := json.Marshal(user)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return
	}
	c.Ctx.ResponseWriter.Write(output)
}
