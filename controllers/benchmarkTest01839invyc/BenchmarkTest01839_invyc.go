//Semgrep original results: []
//Gosec original results: [22]
//CodeQL original results: [22]
//Snyk original results: []
//-------------
//Gosec analysis results: [22, 703]
//CodeQL analysis results: []
//Semgrep analysis results: []
//Snyk analysis results: []
//Original file name: controllers/benchmarkTest01839/BenchmarkTest01839.go
//Original file CWE's: [22]  
//Original file kind: fail
//Mutation info: Insert template from templates-db/languages/go/sensitivity/pointers.tmt with name pointer_init_positive 
//Used extensions: 
//Program:
package controllers

import (
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest01839 struct {
}

func (b *BenchmarkTest01839) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	http.SetCookie(w, &http.Cookie{
		Name:   "BenchmarkTest01839",
		Value:  "FileName",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.URL.Path,
		Domain: r.Host,
	})
	http.ServeFile(w, r, "pathtraver-02/BenchmarkTest01839.html")
}

func (b *BenchmarkTest01839) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookie, err := r.Cookie("BenchmarkTest01839")
	param := "noCookieValueSupplied"
	if err == nil {
		param = cookie.Value

var ptr123 *string = new(string)
*ptr123 = param
param = *ptr123

	}

	bar := doSomething(param)

	var fileName string
	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "TESTFILES_DIR/" + bar
	fos, err = os.Create(fileName)
	if err != nil {
		http.Error(w, "Couldn't open FileOutputStream on file: '"+fileName+"'", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("Now ready to write to file: " + htmlEscape(fileName)))
	if err != nil {
		http.Error(w, "Unable to write response", http.StatusInternalServerError)
	}
}

func doSomething(param string) string {
	num := 106
	if (7*42)-num > 200 {
		return "This should never happen"
	}
	return param
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}

func addSuffix(s *string, suf string) {
	*s = *s + suf
}

func addSuffixDoublePointer(s **string, suf *string) {
	**s = **s + *suf
}

func addSuffixDoublePointerBroken(s **string, suf *string) {
	*s = new(string)
	**s = **s + *suf
}

func swapStrings(a, b *string) {
	temp := *a
	*a = *b
	*b = temp
}

func removeSpaces(s *string) {
    *s = strings.ReplaceAll(*s, " ", "")
}


