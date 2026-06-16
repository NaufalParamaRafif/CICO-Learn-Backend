package domain

import "context"

type UserDialogItem struct {
	ID           uint
	Accuracy     float32
	Feedback     string
	UserID       uint
	DialogItemID uint
}

type UserDialogItemRepository interface {
	GetByID(ctx context.Context, userDialogItemID uint) (UserDialogItem, error)
	GetByUserID(ctx context.Context, userID uint) ([]UserDialogItem, error)
	GetByDialogItemID(ctx context.Context, dialogItemID uint) ([]UserDialogItem, error)
	GetByUserIDAndDialogItemID(ctx context.Context, userID uint, dialogItemID uint) (UserDialogItem, error)

	Insert(ctx context.Context, userDialogItem *UserDialogItem) error
	Update(ctx context.Context, userDialogItem *UserDialogItem) error
	Delete(ctx context.Context, userDialogItemID uint) error
}
