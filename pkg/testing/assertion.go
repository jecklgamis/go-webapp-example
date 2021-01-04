package server

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

// Assert tests if the given condition is true, otherwise fail fatally with a given message
func Assert(t *testing.T, condition bool, message string) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		t.Fatalf("%s:%d: %s", filepath.Base(file), line, message)
	}
}

// Assertf tests if the given condition is true, otherwise fail fatally with a given formatted message similar to Sprintf
func Assertf(t *testing.T, condition bool, format string, args ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		t.Fatalf("%s:%d: %s", filepath.Base(file), line, fmt.Sprintf(format, args...))
	}
}
