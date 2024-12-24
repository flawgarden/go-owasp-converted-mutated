package controllers

import (
"fmt"
"net/http"
"os"
"path/filepath"
)

type BenchmarkTest00362 struct{}

func (b *BenchmarkTest00362) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	param := r.URL.Query().Get("BenchmarkTest00362")
	if param == "" {
		param = ""
	}

	bar := param
	fileName := filepath.Join("testfiles", bar)

stringReturner := func() func() string {
    return func() string {
        return fileName
    }
}

stringRet := stringReturner()
fileName = stringRet()

	var fos *os.File

	defer func() {
		if fos != nil {
			fos.Close()
		}
	}()

	fos, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Couldn't open FileOutputStream on file: '" + fileName + "'")
		return
	}

	_, _ = w.Write([]byte("Now ready to write to file: " + fileName))
}

func main() {
	http.Handle("/pathtraver-00/BenchmarkTest00362", &BenchmarkTest00362{})
	http.ListenAndServe(":8080", nil)
}

func makePrefixer(prefix string) func(string) string {
    return func(value string) string {
        return fmt.Sprintf("%s%s", prefix, value)
    }
}

func makeMessageGenerator(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return prefix + name
		}
	}
}

func makeMessageGeneratorBroken(prefix string) func() func(string) string {
	return func() func(string) string {
		return func(name string) string {
			return name
		}
	}
}


