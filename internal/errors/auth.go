package errs

import "errors"

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUserNotFound = errors.New("user not found")
var ErrUserNameTaken = errors.New("user name is taken")
var ErrUnknownUser = errors.New("unknown user")
