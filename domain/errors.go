package domain

import "errors"

var (
	ErrUserNotFound              = errors.New("User not found")
	ErrEmailExists               = errors.New("Email already exists")
	ErrUsernameTaken             = errors.New("Username has been taken")
	ErrInvalidIdentifierPassword = errors.New("Invalid email, username, or password")

	ErrInvalidFileType  = errors.New("Invalid file type")
	ErrFileSizeTooLarge = errors.New("File size too large")
)
