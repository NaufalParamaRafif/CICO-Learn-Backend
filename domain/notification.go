package domain

import (
	"context"
	"time"
)

type Notification struct {
	Title     string
	Emoticon  string
	Message   string
	SendAt    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	// UsersID []User
}

type NotificationRepository interface {
	GetByID(ctx context.Context, notificationID uint) (Notification, error)
	Insert(ctx context.Context, notification *Notification) error
	Update(ctx context.Context, notification *Notification) error
	Delete(ctx context.Context, notificationID uint) error
}
