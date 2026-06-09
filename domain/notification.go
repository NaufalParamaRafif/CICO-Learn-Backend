package domain

import "time"

type Notification struct {
	Title     string
	Emoticon  string
	Message   string
	SendAt    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	// UsersID []User // TODO: FIX THIS
}
