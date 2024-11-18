package controllers

import (
	"net/http"
	"net/url"
)

type BenchmarkTest00982 struct {
}

func (b *BenchmarkTest00982) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	cookie := http.Cookie{
		Name:   "BenchmarkTest00982",
		Value:  "FOO%3Decho+Injection",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.URL.Path,
		Domain: getDomain(r.URL.String()),
	}
	http.SetCookie(w, &cookie)
	http.ServeFile(w, r, "cmdi-01/BenchmarkTest00982.html")
}

func (b *BenchmarkTest00982) DoPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := r.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest00982" {
			param, _ = url.QueryUnescape(cookie.Value)
			break
		}
	}

	bar := new(Test).DoSomething(r, param)

	cmd := getInsecureOSCommandString()
	args := []string{cmd}
	argsEnv := []string{bar}

	err := execCommand(args, argsEnv, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Test struct {
}

func (t *Test) DoSomething(r *http.Request, param string) string {
	bar := "safe!"
	map20875 := make(map[string]interface{})
	map20875["keyA-20875"] = "a_Value"
	map20875["keyB-20875"] = param
	map20875["keyC"] = "another_Value"
	bar = map20875["keyB-20875"].(string)
	bar = map20875["keyA-20875"].(string)
	return bar
}

func getDomain(url string) string {
	// Вспомогательная функция для получения домена
	// Реализация не представлена
	return ""
}

func getInsecureOSCommandString() string {
	// Вспомогательная функция для получения командной строки
	// Реализация не представлена
	return ""
}

func execCommand(args []string, argsEnv []string, w http.ResponseWriter) error {
	// Вспомогательная функция для выполнения команды
	// Реализация не представлена
	return nil
}
