package domain

import "context"

type TrainWithAITopic struct {
	ID           uint
	Image        string
	ScenarioName string
	Scenario     string
	Rule         string
	Mode         TrainWithAIMode
}

type TrainWithAITopicRepository interface {
	GetByID(ctx context.Context, trainWithAITopicID uint) (TrainWithAITopic, error)
	Insert(ctx context.Context, trainWithAITopic *TrainWithAITopic) error
	Update(ctx context.Context, trainWithAITopic *TrainWithAITopic) error
	Delete(ctx context.Context, trainWithAITopicID uint) error
}
