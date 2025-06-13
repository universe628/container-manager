package errs

import "errors"

var ErrDeadlineExceeded = errors.New("deadline exceeded")
var ErrCancelled = errors.New("request cancelled")
