package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type BenchmarkTest01876 struct {
}

func (bt *BenchmarkTest01876) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	userCookie := http.Cookie{
		Name:   "BenchmarkTest01876",
		Value:  "my_userid",
		MaxAge: 60 * 3,
		Secure: true,
		Path:   r.URL.Path,
		Domain: r.Host,
	}
	http.SetCookie(w, &userCookie)
	http.ServeFile(w, r, "trustbound-01/BenchmarkTest01876.html")
}

func (bt *BenchmarkTest01876) HandlePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	cookies := r.Cookies()
	param := "noCookieValueSupplied"
	for _, cookie := range cookies {
		if cookie.Name == "BenchmarkTest01876" {
			param = cookie.Value
			break
		}
	}

	bar := doSomething(param)

	r.AddCookie(&http.Cookie{Name: "userid", Value: bar})

	w.Write([]byte("Item: 'userid' with value: '" + bar + "' saved in session."))
}

func doSomething(param string) string {
	bar := param
	if param != "" && len(param) > 1 {
		bar = param[:len(param)-1]
	}
	return bar
}

func TestBenchmarkTest01876(t *testing.T) {
	req := httptest.NewRequest("GET", "/trustbound-01/BenchmarkTest01876", nil)
	w := httptest.NewRecorder()

	bt := &BenchmarkTest01876{}
	bt.ServeHTTP(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK; got %v", res.Status)
	}

	// Additional tests for the POST method can be added here
}
