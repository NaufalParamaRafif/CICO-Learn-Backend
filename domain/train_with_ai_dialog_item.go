package domain

import "context"

type TrainWithAIDialogItem struct {
	ID                  uint   `db:"id"`
	OrderDialog         uint16 `db:"order_dialog"`
	Dialog              string `db:"dialog"`
	IsFromUser          bool   `db:"is_from_user"`
	TrainWithAIDialogID uint   `db:"train_with_ai_dialog_id"`
}

type TrainWithAIDialogItemRepository interface {
	GetByID(ctx context.Context, trainWithAIDialogID uint) (TrainWithAIDialogItem, error)
	Insert(ctx context.Context, trainWithAIDialog *TrainWithAIDialogItem) error
	Update(ctx context.Context, trainWithAIDialog *TrainWithAIDialogItem) error
	Delete(ctx context.Context, trainWithAIDialogID uint) error
}
