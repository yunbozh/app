package main

import (
	"app/common/logger"
	"runtime/debug"
)

func Recover() {
	if err := recover(); err != nil {
		logger.Error("%v", err)
		logger.Error("%s", debug.Stack())
	}
}
