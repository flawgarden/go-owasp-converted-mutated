package controllers

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"

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

type BenchmarkTest00635 struct {
	beego.Controller
}

func (c *BenchmarkTest00635) Get() {
	c.Post()
}

func (c *BenchmarkTest00635) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := c.GetString("BenchmarkTest00635")
	if param == "" {
		param = ""
	}

	bar := param // Assume some encoding function applied here

	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	md := sha1.New()
	input := []byte(bar)
	md.Write(input)
	result := md.Sum(nil)

	c.SaveHash(result)

	output := fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", bar)
	c.Ctx.ResponseWriter.Write([]byte(output))
}

func (c *BenchmarkTest00635) SaveHash(hash []byte) {
	fileTarget := "passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	fw.WriteString(fmt.Sprintf("hash_value=%s\n", base64.StdEncoding.EncodeToString(hash)))
}
