package util

import "github.com/brieefly/log"

// ErrorHandler - an interface allowing custom
type ErrorHandler interface {
	HandleError()
}

func HandleError(customHandler func(error) error) {
	defer func() {
		if r := recover(); r != nil {
			log.Error(r)
		} else {

		}
	}()
}
