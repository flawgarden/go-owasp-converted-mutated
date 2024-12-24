package controllers

import (
"fmt"
"net/http"
"os"
"strings"
"sync"
)

type BenchmarkTest02205 struct{}

func (bt *BenchmarkTest02205) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	param := r.FormValue("BenchmarkTest02205")

	bar := doSomething(param)

	fileName := ""
	var fos *os.File

	defer func() {

w := &Wrapper[string]{Value: bar}
task1 := NewSwitchingTask(w, bar)
task2 := NewSwitchingTask(w, bar)
var wg sync.WaitGroup
wg.Add(2)
go func() {
    defer wg.Done()
    task2.Run()
}()
go func() {
    defer wg.Done()
    task1.Run()
}()
wg.Wait()
bar = w.Value

		if fos != nil {
			fos.Close()
		}
	}()

	fileName = "testfiles/" + bar

	fos, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	response := fmt.Sprintf("Now ready to write to file: %s", htmlspecialchars(fileName))
	w.Write([]byte(response))
}

func doSomething(param string) string {
	var bar string
	num := 196
	if (500/42)+num > 200 {
		bar = param
	} else {
		bar = "This should never happen"
	}
	return bar
}

func htmlspecialchars(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "<", "&lt;")
}
