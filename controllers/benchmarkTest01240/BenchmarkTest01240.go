package controllers

import (
	"fmt"
	"net/http"
	"os"
)

type BenchmarkTest01240 struct{}

func (b *BenchmarkTest01240) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01240")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	fileName := ""
	fos := new(os.File)

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fileName = fmt.Sprintf("testfiles/%s", bar)

	var err error
	fos, err = os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	w.Write([]byte(fmt.Sprintf("Now ready to write to file: %s", fileName)))
}

func (b *BenchmarkTest01240) doSomething(param string) string {
	bar := "alsosafe"
	if param != "" {
		valuesList := []string{"safe", param, "moresafe"}
		valuesList = valuesList[1:] // remove the 1st safe value
		bar = valuesList[0]         // get the last 'safe' value
	}
	return bar
}

func main() {
	http.Handle("/", &BenchmarkTest01240{})
	http.ListenAndServe(":8080", nil)
}
