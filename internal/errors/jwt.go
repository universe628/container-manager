package errs

import "errors"

var ErrUnknownToken = errors.New("token type unknown")
var ErrTokenInvalid = errors.New("authorization token invalid")
