package serror

import (
	"fmt"
	"runtime/debug"
)

// Serror represents a structured error.
type Serror interface {
	GetError() error //gets error
}

// serror is the implementation of the Serror interface.
type serror struct {
	err   error
	stack []byte
	code  int
}

// GetError returns the underlying error.
func (s *serror) GetError() error {
	return s.err
}

// NewSerror creates a new instance of Serror.
func NewSerror(err error, code int) Serror {
	fmt.Printf("Error: %v\nStack Trace:\n%s\n", err, debug.Stack()) //log error with stacktrace during conversion
	return &serror{
		err:   err,
		stack: debug.Stack(),
		code:  code,
	}
}
