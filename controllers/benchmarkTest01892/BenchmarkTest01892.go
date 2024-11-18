package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type BenchmarkTest01892 struct {
}

func (b *BenchmarkTest01892) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest01892",
		Value:  "2222",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.URL.Path,
		Domain: getDomain(r.URL.String()),
	}
	http.SetCookie(w, &userCookie)
	http.ServeFile(w, r, "xpathi-00/BenchmarkTest01892.html")
}

func (b *BenchmarkTest01892) doPost(r *http.Request, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := r.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01892" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(r, param)

	filePath := filepath.Join("classpath", "employees.xml")
	xmlDocument, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintln(w, "Error reading XML file")
		return
	}

	expression := fmt.Sprintf("/Employees/Employee[@emplid='%s']", bar)
	nodeList := evaluateXPath(xmlDocument, expression)

	w.Write([]byte("Your query results are: <br/>"))
	for _, value := range nodeList {
		w.Write([]byte(value + "<br/>"))
	}
}

func doSomething(r *http.Request, param string) string {
	bar := "safe!"
	map3451 := make(map[string]interface{})
	map3451["keyA-3451"] = "a-Value"
	map3451["keyB-3451"] = param
	map3451["keyC"] = "another-Value"
	bar = map3451["keyB-3451"].(string)
	return bar
}

func getDomain(urlStr string) string {
	u, _ := url.Parse(urlStr)
	return u.Hostname()
}

func evaluateXPath(xmlDocument []byte, expression string) []string {
	// Implement the logic to evaluate XPath expression against xmlDocument
	// Return the matched values
	return nil // Placeholder return
}
