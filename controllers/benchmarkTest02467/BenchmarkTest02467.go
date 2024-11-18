package controllers

import (
	"fmt"
	"net/http"
	"os"
)

func BenchmarkTest02467(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	values := r.Form["BenchmarkTest02467"]
	var param string
	if len(values) > 0 {
		param = values[0]
	} else {
		param = ""
	}

	bar := doSomething(param)

	var fileName string
	var file *os.File

	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	fileName = "testfiles/" + bar
	var err error
	file, err = os.Open(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileInputStream on file: '" + fileName + "'")
		return
	}

	b := make([]byte, 1000)
	size, err := file.Read(b)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	fmt.Fprintf(w, "The beginning of file: '%s' is:\n\n", fileName)
	w.Write([]byte(htmlEscape(string(b[:size]))))
}

func doSomething(param string) string {
	a := param
	b := a + " SafeStuff"
	b = b[:len(b)-len("Chars")] + "Chars"
	c := map[string]interface{}{"key": b}
	d := c["key"].(string)[:len(c["key"].(string))-1]
	e := string(decodeBase64(encodeBase64([]byte(d))))
	f := e[:len(e)-7]
	return f
}

func htmlEscape(str string) string {
	return str // Здесь должна быть логика для экранирования строки
}

func encodeBase64(data []byte) []byte {
	return data // Здесь должна быть логика для кодирования в Base64
}

func decodeBase64(data []byte) []byte {
	return data // Здесь должна быть логика для декодирования из Base64
}

func main() {
	http.HandleFunc("/pathtraver-03/BenchmarkTest02467", BenchmarkTest02467)
	http.ListenAndServe(":8080", nil)
}
