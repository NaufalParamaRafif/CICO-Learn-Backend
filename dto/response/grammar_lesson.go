package response

type GrammarLessonSection struct {
	Title        string   `json:"title"`
	SectionOrder uint8    `json:"section_order"`
	Explanations []string `json:"explanations"`
}

type GrammarLessonExercise struct {
	Question      string `json:"question"`
	AChoice       string `json:"a_choice"`
	BChoice       string `json:"b_choice"`
	CChoice       string `json:"c_choice"`
	DChoice       string `json:"d_choice"`
	CorrectChoice string `json:"correct_choice"`
	Explanation   string `json:"explanation"`
}

type GrammarLessonExplanationResponse struct {
	Text  string `json:"text"`
	Order uint8  `json:"order"`
	Type  string `json:"type"`
}

type GrammarLessonGetStatusOKResponseResponse struct {
	Title    string                 `json:"title"`
	Meaning  string                 `json:"meaning"`
	Sections []GrammarLessonSection `json:"sections"`
	Exercise string                 `json:"exercise"`
}

type GrammarLessonExercisesGetStatusOKResponseResponse struct {
	Exercises []GrammarLessonExercise `json:"exercises"`
}

type GrammarLessonExercisesPostStatusOKResponseResponse struct {
	OverallScore float32 `json:"overall_score"`
}
