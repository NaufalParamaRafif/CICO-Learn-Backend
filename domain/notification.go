package domain

import (
	"context"
	"time"
)

type Notification struct {
	Title     string    `db:"title"`
	Emoticon  string    `db:"emoticon"`
	Message   string    `db:"message"`
	SendAt    time.Time `db:"send_at"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	// UsersID []User
}

type NotificationRepository interface {
	GetByID(ctx context.Context, notificationID uint) (Notification, error)
	Insert(ctx context.Context, notification *Notification) error
	Update(ctx context.Context, notification *Notification) error
	Delete(ctx context.Context, notificationID uint) error
}
