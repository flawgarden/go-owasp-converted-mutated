package controllers

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type SqlInjectionVuln2Controller struct {
	web.Controller
}

func init() {
	web.Router("/ldapi-00/BenchmarkTest01743", &SqlInjectionVuln2Controller{})
}

func (c *SqlInjectionVuln2Controller) Get() {
	c.Post()
}

func (c *SqlInjectionVuln2Controller) Post() {
	response := c.Ctx.ResponseWriter
	response.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := c.GetString("BenchmarkTest01743")
	bar := decode(param)

	db, err := sql.Open("mysql", source)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	filter := fmt.Sprintf("(&(objectclass=person))(|(uid=%s)(street={0}))", bar)
	results, err := db.Query("SELECT uid, street FROM user WHERE id = ?", strings.TrimSpace(bar))
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	defer results.Close()

	found := false
	for results.Next() {
		var uid, street string
		if err := results.Scan(&uid, &street); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Write([]byte(fmt.Sprintf("LDAP query results:<br>Record found with name %s<br>Address: %s<br>", uid, street)))
		found = true
	}

	if !found {
		response.Write([]byte(fmt.Sprintf("LDAP query results: nothing found for query: %s", filter)))
	}
}

func decode(param string) string {
	if param == "" {
		return ""
	}
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(param)))
	n, _ := base64.StdEncoding.Decode(decoded, []byte(param))
	return string(decoded[:n])
}
