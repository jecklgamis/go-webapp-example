package server

import (
	"runtime"
	"testing"
)

func Assert(t *testing.T, condition bool, message string) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s (%s:%d)", message, file, line)
	}
}
