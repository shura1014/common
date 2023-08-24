package common

import (
	"net/http"
	"testing"
	"time"
)

func TestDeferredResult_SetDeferredResultHandler(t *testing.T) {
	result := NewDeferredResult(5*time.Second, "超时了")
	result.SetDeferredResultHandler(func(result any) {
		t.Log(result)
	})
	time.Sleep(3 * time.Second)
	result.SetResult("success")
}

type testHandler struct {
}

func (test *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	NewDeferredResultWriter(5*time.Second, "超时了", w)
}

func TestNewHttpDeferredResult(t *testing.T) {
	http.Handle("/", &testHandler{})
	_ = http.ListenAndServe(":8000", nil)
}
