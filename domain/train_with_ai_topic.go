package domain

// TODO: FIX MIGRATION

type TrainWithAITopic struct {
	ID           uint
	Image        string
	ScenarioName string
	Scenario     string
	Mode         TrainWithAIMode
	Rule         string
}
