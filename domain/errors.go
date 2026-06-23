package domain

import "errors"

var (
	ErrUserNotFound              = errors.New("User not found")
	ErrEmailExists               = errors.New("Email already exists")
	ErrUsernameTaken             = errors.New("Username has been taken")
	ErrInvalidIdentifierPassword = errors.New("Invalid email, username, or password")

	ErrInvalidFileType  = errors.New("Invalid file type")
	ErrFileSizeTooLarge = errors.New("File size too large")

	ErrEmptyPassword       = errors.New("Password must be filled")
	ErrEmptyIdentifier     = errors.New("Email or username must be filled")
	ErrEnglishLevelInvalid = errors.New("English level is invalid")
	ErrLanguageInvalid     = errors.New("Language is invalid")
	ErrTargetInvalid       = errors.New("Target is invalid")
	ErrTimeTargetInvalid   = errors.New("Time Target is invalid")
	ErrGenderInvalid       = errors.New("Gender is invalid")
)
