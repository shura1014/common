package common

import (
	"encoding/json"
	"github.com/shura1014/common/clog"
	"github.com/shura1014/common/goerr"
	"io"
	"time"
)

type DeferredResult struct {
	timeout       time.Duration
	timeoutResult any
	resultHandler func(result any)
	done          bool
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

func NewDeferredResult(timeout time.Duration, timeoutResult any, callback ...func(deferred *DeferredResult)) {
	instance := DeferredResultInstance(timeout, timeoutResult)
	if len(callback) > 0 {
		go callback[0](instance)
	}
	instance.listener()
}

func (deferred *DeferredResult) SetDeferredResultHandler(resultHandler func(result any)) {
	deferred.resultHandler = resultHandler
}

func (deferred *DeferredResult) SetResult(result any) error {
	if deferred.done {
		return goerr.Text("the deferred result is done")
	}
	deferred.result <- result
	return nil
}

func (deferred *DeferredResult) listener() {
	defer func() { deferred.done = true }()
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

func NewDeferredResultWriter(timeout time.Duration, timeoutResult any, write io.Writer, callback ...func(deferred *DeferredResultWriter)) {
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
	if len(callback) > 0 {
		go callback[0](deferred)
	}
	deferred.listener()
}
