package constants

import "errors"

var (
	ErrUserAlreadyExist = errors.New("User already exist")
	ErrUserNotFound     = errors.New("User is not found")
	ErrInternalServer   = errors.New("Unexpected error happened. Please try again!")
)
