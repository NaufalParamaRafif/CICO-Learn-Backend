package dto

type ActivityThisMonth struct {
	Day                uint8  `json:"day"`
	TotalLessonLearned uint16 `json:"total_lesson_learned"`
}

type UserGetStatusOKResponse struct {
	BackgroundImage   string              `json:"background_image"`
	ProfileImage      string              `json:"profile_image"`
	FullName          string              `json:"full_name"`
	ActivityThisMonth []ActivityThisMonth `json:"activity_this_month"`
	Level             EnglishLevel        `json:"level"`
	AppLanguage       Language            `json:"app_language"`
	Target            Language            `json:"target"`
	DailyReminder     []string            `json:"daily_reminder"` // format: HH:MM
	TimeTarget        TimeTarget          `json:"time_target"`
}
