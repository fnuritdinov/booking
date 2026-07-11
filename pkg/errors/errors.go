package errors

import "errors"

var ErrValidate = errors.New("err from validate")
var ErrBadRequest = errors.New("bad request")
var ErrNotFound = errors.New("not found")
