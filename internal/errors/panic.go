package errs

import (
	"errors"
	"fmt"
)

var ErrPanic = errors.New("panic occurred")

type PanicError struct {
	Message string
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("%s: %s", ErrPanic.Error(), e.Message)
}
