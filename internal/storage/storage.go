package storage

import "errors"

var (
	ErrUserExists   = errors.New("the User already exists")
	ErrUserNotFound = errors.New("the User was not found")
	ErrAppNotFound  = errors.New("the app was not found")
)
