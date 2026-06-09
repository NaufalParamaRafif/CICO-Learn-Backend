package domain

import "time"

type TimeOfDay struct {
	Hour   uint8
	Minute uint8
}

type User struct {
	ID              uint
	FirstName       string
	LastName        string
	Username        string
	Email           string
	Address         string
	Gender          Gender
	ProfileImage    string
	BackgroundImage string
	Level           EnglishLevel
	Language        Language
	Target          Target
	DailyReminder   []TimeOfDay
	TimeTarget      TimeTarget
	Password        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
