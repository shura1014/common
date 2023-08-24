package common

import (
	"net/http"
	"testing"
	"time"
)

func TestDeferredResult_SetDeferredResultHandler(t *testing.T) {
	var result *DeferredResult

	NewDeferredResult(5*time.Second, "超时了", func(deferred *DeferredResult) {
		result = deferred
		result.SetDeferredResultHandler(func(result any) {
			t.Log(result)
		})
		//time.Sleep(3 * time.Second)
		_ = result.SetResult("success")
	})

}

type testHandler struct {
}

func (test *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var deferredResultWriter *DeferredResultWriter
	NewDeferredResultWriter(5*time.Second, "超时了", w, func(deferred *DeferredResultWriter) {
		deferredResultWriter = deferred
		time.Sleep(3 * time.Second)
		_ = deferredResultWriter.SetResult("success")
	})

}

func TestNewHttpDeferredResult(t *testing.T) {
	http.Handle("/", &testHandler{})
	_ = http.ListenAndServe(":8000", nil)
}
