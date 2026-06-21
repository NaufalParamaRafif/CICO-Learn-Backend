package domain

import "context"

type TrainWithAITopic struct {
	ID           uint            `db:"id"`
	Image        string          `db:"image"`
	ScenarioName string          `db:"scenario_name"`
	Scenario     string          `db:"scenario"`
	Rule         string          `db:"rule"`
	Mode         TrainWithAIMode `db:"mode"`
}

type TrainWithAITopicRepository interface {
	GetByID(ctx context.Context, trainWithAITopicID uint) (TrainWithAITopic, error)
	Insert(ctx context.Context, trainWithAITopic *TrainWithAITopic) error
	Update(ctx context.Context, trainWithAITopic *TrainWithAITopic) error
	Delete(ctx context.Context, trainWithAITopicID uint) error
}
