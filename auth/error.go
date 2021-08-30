package auth

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidAccesToken = errors.New("invalid access token")
)
