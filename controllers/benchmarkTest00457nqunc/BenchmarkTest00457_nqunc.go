//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [22, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest00457/BenchmarkTest00457.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/collections/map.tmt with name map_merge_2_negative 
//Used extensions: MACRO_Create_Map -> ~[MACRO_MapName]~ := make(map[~[TYPE@1]~]~[TYPE@2]~) | MACRO_Add_EXPR_ToMap -> ~[MACRO_MapName]~[~[EXPR_~[TYPE@1]~@1]~] = ~[EXPR_~[TYPE@2]~@2]~ | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234 | MACRO_MapName -> map787234
//Program:
package controllers

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"path/filepath"
)

type BenchmarkTest00457 struct{}

func (b *BenchmarkTest00457) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest00457) doPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	param := r.FormValue("BenchmarkTest00457")

	bar := ""
	if param != "" {

map787234 := make(map[string]string)
map787234["ETNfb"] = "QOiTd"
map787234["ETNfb"] = bar
param = map787234["ETNfb"]
map787234["ETNfb"] += bar

		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the param value
	}

	fileName := filepath.Join("testfiles", bar)
	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}
	defer fos.Close()

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.Write([]byte("Now ready to write to file: " + html.EscapeString(fileName)))
}
