package domain

import (
	"context"
	"time"
)

type TrainWithAIDialog struct {
	ID                 uint
	Mode               TrainWithAIMode
	UserID             uint
	TrainWithAITopicID uint
	StartedAt          time.Time
	LastOpenedAt       time.Time
}

type TrainWithAIDIalogRepository interface {
	GetByID(ctx context.Context, trainWithAIDialogID uint) (TrainWithAIDialog, error)
	GetByUserID(ctx context.Context, userID uint) ([]TrainWithAIDialog, error)
	GetByTrainWithAITopicID(ctx context.Context, trainWithAITopicID uint) ([]TrainWithAIDialog, error)
	GetByUserIDAndTrainWithAITopicID(ctx context.Context, userID uint, trainWithAITopicID uint) (TrainWithAIDialog, error)

	Insert(ctx context.Context, trainWithAIDialog *TrainWithAIDialog) error
	Update(ctx context.Context, trainWithAIDialog *TrainWithAIDialog) error
	Delete(ctx context.Context, trainWithAIDialogID uint) error

	GetTrainWithAIDialogItems(ctx context.Context, trainWithAIDialogID uint) ([]TrainWithAIDialogItem, error)
}
