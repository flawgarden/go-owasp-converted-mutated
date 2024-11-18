package controllers

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const source = "root:password@tcp(127.0.0.1:3306)/goseccode"

type BenchmarkTest02392 struct {
}

func (bt *BenchmarkTest02392) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This input source requires a POST", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	param := r.FormValue("BenchmarkTest02392")
	if param == "" {
		param = ""
	}
	bar := doSomething(param)

	algorithm := "SHA512"
	hashValue := hashInput(bar, algorithm)

	fileTarget, err := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Unable to open file", http.StatusInternalServerError)
		return
	}
	defer fileTarget.Close()
	_, err = fileTarget.WriteString(fmt.Sprintf("hash_value=%s\n", hashValue))
	if err != nil {
		http.Error(w, "Unable to write to file", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", encodeForHTML(bar))
	fmt.Fprintln(w, "Hash Test executed")
}

func doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func hashInput(input string, algorithm string) string {
	// Implementation for hashing would go here
	return "hashed_value" // Placeholder return
}

func encodeForHTML(input string) string {
	// HTML encoding implementation would go here
	return input // Placeholder return
}
