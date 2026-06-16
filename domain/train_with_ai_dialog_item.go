package domain

import "context"

type TrainWithAIDialogItem struct {
	ID                  uint
	OrderDialog         uint16
	Dialog              string
	IsFromUser          bool
	TrainWithAIDialogID uint
}

type TrainWithAIDialogItemRepository interface {
	GetByID(ctx context.Context, trainWithAIDialogID uint) (TrainWithAIDialogItem, error)
	Insert(ctx context.Context, trainWithAIDialog *TrainWithAIDialogItem) error
	Update(ctx context.Context, trainWithAIDialog *TrainWithAIDialogItem) error
	Delete(ctx context.Context, trainWithAIDialogID uint) error
}
