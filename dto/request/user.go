package request

type ActivityThisMonth struct {
	Day                uint8  `json:"day"`
	TotalLessonLearned uint16 `json:"total_lesson_learned"`
}

type UserRegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateProfileRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
}

type UpdateLearningPreferenceRequest struct {
	Level         string   `json:"level"`
	Language      string   `json:"language"`
	Target        string   `json:"target"`
	DailyReminder []string `json:"daily_reminder"`
	TimeTarget    string   `json:"time_target"`
}
