package response

type VocabularyComponent struct {
	LessonType      string  `json:"lesson_type"`
	Image           string  `json:"image"`
	Title           string  `json:"title"`
	TotalVocabulary uint16  `json:"total_vocabulary"`
	Type            string  `json:"type"`
	Progress        float32 `json:"progress"`
}

type GrammarComponent struct {
	LessonType    string  `json:"lesson_type"`
	Image         string  `json:"image"`
	Title         string  `json:"title"`
	TotalSection  uint8   `json:"total_section"`
	TotalExercise uint16  `json:"total_exercise"`
	Progress      float32 `json:"progress"`
}

type DialogComponent struct {
	LessonType  string  `json:"lesson_type"`
	Image       string  `json:"image"`
	Title       string  `json:"title"`
	TotalDialog uint16  `json:"total_dialog"`
	Progress    float32 `json:"progress"`
}
