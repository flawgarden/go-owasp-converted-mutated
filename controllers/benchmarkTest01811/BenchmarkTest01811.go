package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest01811Controller struct {
	web.Controller
}

func (c *BenchmarkTest01811Controller) Get() {
	c.Post()
}

func (c *BenchmarkTest01811Controller) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01811")
	bar := doSomething(param)

	sqlStr := fmt.Sprintf("SELECT TOP 1 userid FROM USERS WHERE USERNAME='foo' AND PASSWORD='%s'", bar)
	results, err := queryForMap(sqlStr)
	if err != nil {
		c.Ctx.ResponseWriter.Write([]byte("Error processing request."))
		return
	}
	output, _ := json.Marshal(results)
	c.Ctx.ResponseWriter.Write(output)
}

func doSomething(param string) string {
	bar := "safe!"
	map31047 := make(map[string]interface{})
	map31047["keyA-31047"] = "a_Value"
	map31047["keyB-31047"] = param
	map31047["keyC"] = "another_Value"
	bar = map31047["keyB-31047"].(string)
	bar = map31047["keyA-31047"].(string)
	return bar
}

func queryForMap(sqlStr string) (map[string]interface{}, error) {
	db, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var results map[string]interface{}
	err = db.QueryRow(sqlStr).Scan(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
