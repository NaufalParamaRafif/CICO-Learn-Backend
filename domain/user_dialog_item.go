package domain

import "context"

type UserDialogItem struct {
	ID           uint    `db:"id"`
	Accuracy     float32 `db:"accuracy"`
	Feedback     string  `db:"feedback"`
	UserID       uint    `db:"user_id"`
	DialogItemID uint    `db:"dialog_item_id"`
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
