package domain

import (
	"context"
	"time"
)

type UserNotification struct {
	ID             uint `db:"id"`
	UserID         uint `db:"user_id"`
	NotificationID uint `db:"notification_id"`
	ReadAt         time.Time `db:"read_at"`
}

type UserNotificationRepository interface {
	GetByID(ctx context.Context, userNotificationID uint) (UserNotification, error)
	GetByUserID(ctx context.Context, userID uint) ([]UserNotification, error)
	GetByNotificationID(ctx context.Context, notificationID uint) ([]UserNotification, error)
	GetByUserIDAndNotificationID(ctx context.Context, userID uint, notificationID uint) (UserNotification, error)

	Insert(ctx context.Context, userNotification *UserNotification) error
	Update(ctx context.Context, userNotification *UserNotification) error
	Delete(ctx context.Context, userNotificationID uint) error
}
