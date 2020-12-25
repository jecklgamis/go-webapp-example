package server

import (
	"fmt"
	"runtime"
	"testing"
)

func Assert(t *testing.T, condition bool, message string) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf(fmt.Sprintf("%s (%s:%d)", message, file, line))
	}
}
