package appErr

import "errors"

var (
	NotFound         = errors.New("user not found")
	Unauthorized     = errors.New("Unauthorized")
	MethodNotAllowed = errors.New("Unauthorized")
	Internal         = errors.New("Internal Error")
)
