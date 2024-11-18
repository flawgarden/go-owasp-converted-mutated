package controllers

import (
	"fmt"
	"html"
	"net/http"
	"net/url"

	"xorm.io/xorm"
)

type BenchmarkTest01909 struct {
	DB *xorm.Engine
}

func (b *BenchmarkTest01909) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	param := r.Header.Get("BenchmarkTest01909")
	param, _ = url.QueryUnescape(param)

	bar := doSomething(r, param)

	filter := fmt.Sprintf("(&(objectclass=person)(uid=%s))", bar)
	results, err := b.DB.Query("SELECT * FROM ldap WHERE filter = ?", filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	found := false
	for _, result := range results {
		w.Write([]byte(fmt.Sprintf("LDAP query results:<br>Record found with name %s<br>Address: %s<br>",
			result["uid"], result["street"])))
		found = true
	}

	if !found {
		w.Write([]byte(fmt.Sprintf("LDAP query results: nothing found for query: %s", html.EscapeString(filter))))
	}
}

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map30748 := make(map[string]interface{})
	map30748["keyA-30748"] = "a_Value"
	map30748["keyB-30748"] = param
	map30748["keyC"] = "another_Value"
	bar = map30748["keyB-30748"].(string)
	bar = map30748["keyA-30748"].(string)

	return bar
}
