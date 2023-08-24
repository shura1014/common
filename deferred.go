package common

import (
	"encoding/json"
	"github.com/shura1014/common/clog"
	"io"
	"time"
)

type DeferredResult struct {
	timeout       time.Duration
	timeoutResult any
	resultHandler func(result any)
	result        chan any
}

func DeferredResultInstance(timeout time.Duration, timeoutResult any) *DeferredResult {
	deferred := &DeferredResult{
		timeout:       timeout,
		timeoutResult: timeoutResult,
		result:        make(chan any),
	}
	return deferred
}

func NewDeferredResult(timeout time.Duration, timeoutResult any) *DeferredResult {
	instance := DeferredResultInstance(timeout, timeoutResult)
	go instance.listener()
	return instance
}

func (deferred *DeferredResult) SetDeferredResultHandler(resultHandler func(result any)) {
	deferred.resultHandler = resultHandler
}

func (deferred *DeferredResult) SetResult(result any) {
	deferred.result <- result
}

func (deferred *DeferredResult) listener() {
	timer := time.NewTimer(deferred.timeout)
	select {
	case result := <-deferred.result:
		deferred.resultHandler(result)
	case <-timer.C:
		deferred.resultHandler(deferred.timeoutResult)
	}
}

type DeferredResultWriter struct {
	*DeferredResult
	write io.Writer
}

func NewDeferredResultWriter(timeout time.Duration, timeoutResult any, write io.Writer) *DeferredResultWriter {
	deferred := &DeferredResultWriter{
		DeferredResult: DeferredResultInstance(timeout, timeoutResult),
		write:          write,
	}
	deferred.SetDeferredResultHandler(func(result any) {
		marshal, err := json.Marshal(result)
		if err != nil {
			clog.Error(err)
		}
		_, err = deferred.write.Write(marshal)
		if err != nil {
			clog.Error(err)
		}
	})
	deferred.listener()
	return deferred
}
