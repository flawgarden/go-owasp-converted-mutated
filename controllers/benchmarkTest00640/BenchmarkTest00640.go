package controllers

import (
	"bufio"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"html/template"
	"net/http"
	"os"
	"strings"

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

type BenchmarkTest00640 struct {
	beego.Controller
}

func (c *BenchmarkTest00640) Get() {
	c.Post()
}

func (c *BenchmarkTest00640) Post() {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest00640")
	if param == "" {
		param = ""
	}

	bar := ""
	num := 106
	if (7*42)-num > 200 {
		bar = "This should never happen"
	} else {
		bar = param
	}

	algorithm := "SHA5"
	if value, ok := getBenchmarkProps("hashAlg2"); ok {
		algorithm = value
	}

	hashValue, err := hashValue(algorithm, bar)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, "Hashing error", http.StatusInternalServerError)
		return
	}

	if err := storeHashValue(hashValue); err != nil {
		http.Error(c.Ctx.ResponseWriter, "Unable to store hash value", http.StatusInternalServerError)
		return
	}

	c.Ctx.WriteString(fmt.Sprintf("Sensitive value '%s' hashed and stored<br/>", htmlEncode(bar)))
	c.Ctx.WriteString("Hash Test executed")
}

func getBenchmarkProps(key string) (string, bool) {
	benchmarkprops := make(map[string]string)
	file, err := os.Open("benchmark.properties")
	if err != nil {
		return "", false
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		if len(line) == 2 {
			benchmarkprops[line[0]] = line[1]
		}
	}
	return benchmarkprops[key], true
}

func hashValue(algorithm, input string) ([]byte, error) {
	md, err := getMessageDigest(algorithm)
	if err != nil {
		return nil, err
	}
	md.Write([]byte(input))
	return md.Sum(nil), nil
}

func getMessageDigest(algorithm string) (hash.Hash, error) {
	switch algorithm {
	case "SHA1":
		return sha1.New(), nil
	case "SHA256":
		return sha256.New(), nil
	default:
		return nil, fmt.Errorf("unsupported hash algorithm: %s", algorithm)
	}
}

func storeHashValue(hashValue []byte) error {
	fileTarget := "path/to/file/passwordFile.txt"
	fw, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fw.Close()
	_, err = fw.WriteString(fmt.Sprintf("hash_value=%s\n", base64.StdEncoding.EncodeToString(hashValue)))
	return err
}

func htmlEncode(input string) string {
	return template.HTMLEscapeString(input)
}
