package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type BenchmarkTest01248 struct {
}

func (b *BenchmarkTest01248) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.doPost(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (b *BenchmarkTest01248) doPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	param := r.URL.Query().Get("BenchmarkTest01248")
	if param == "" {
		param = ""
	}

	bar := b.doSomething(param)

	benchmarkProps, _ := os.ReadFile("benchmark.properties")
	props := make(map[string]string)
	for _, line := range strings.Split(string(benchmarkProps), "\n") {
		if line != "" {
			kv := strings.SplitN(line, "=", 2)
			if len(kv) == 2 {
				props[kv[0]] = kv[1]
			}
		}
	}

	algorithm := props["hashAlg1"]
	if algorithm == "" {
		algorithm = "SHA512"
	}

	// Example hashing logic; replace with actual implementation
	hash := fmt.Sprintf("%x", bar) // Placeholder for actual hash computation

	file, _ := os.OpenFile("passwordFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(fmt.Sprintf("hash_value=%s\n", hash))

	fmt.Fprintf(w, "Sensitive value '%s' hashed and stored<br/>", htmlEscape(bar))
	fmt.Fprintln(w, "Hash Test executed")
}

func (b *BenchmarkTest01248) doSomething(param string) string {
	bar := param
	if len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func htmlEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "<", "&lt;"), ">", "&gt;")
}
