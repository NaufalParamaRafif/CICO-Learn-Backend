package dto

type TrainWithAITopic struct {
	Image    string `json:"image"`
	Title    string `json:"title"`
	Scenario string `json:"scenario"`
}

type TrainWithAIHistory struct {
	Mode              TrainWithAIMode `json:"mode"`
	Image             string          `json:"image"`
	Topic             string          `json:"topic"`
	TotalTrainingTime string          `json:"total_training_time"` // FORMAT: 1 Minute, 30 Minutes, 1 Hour, 2 Hours
	LastOpened        string          `json:"last_opened"`         // FORMAT: TIME WITH TIMEZONE
}

type TrainWithAITopicsGetStatusOKResponse struct {
	Topics []TrainWithAITopic `json:"topics"`
}

type TrainWithAIHistoriesGetStatusOKResponse struct {
	Histories []TrainWithAIHistory `json:"histories"`
}
