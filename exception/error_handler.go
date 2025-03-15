package exception

import "errors"

var (
	ErrNotFound       = errors.New("data not found")
	ErrInternalServer = errors.New("internal server error")
)
