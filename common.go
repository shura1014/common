package common

import "github.com/shura1014/common/goerr"

func TryCatch(try func(), catch ...func(exception error)) {
	defer func() {
		if exception := recover(); exception != nil && len(catch) > 0 {
			if v, ok := exception.(error); ok && goerr.IsStack(v) {
				catch[0](v)
			} else {
				catch[0](goerr.Wrap(exception))
			}
		}
	}()
	try()
}
