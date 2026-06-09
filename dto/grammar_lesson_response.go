package dto

import "github.com/NaufalParamaRafif/CICO-Learn-Backend/domain"

type GrammarLessonExplanation struct {
	Text  string                              `json:"text"`
	Order uint8                               `json:"order"`
	Type  domain.GrammarLessonExplanationType `json:"type"`
}

type GrammarLessonSection struct {
	Title        string                     `json:"title"`
	SectionOrder uint8                      `json:"section_order"`
	Explanations []GrammarLessonExplanation `json:"explanations"`
}

type GrammarLessonExercise struct {
	Question      string                             `json:"question"`
	AChoice       string                             `json:"a_choice"`
	BChoice       string                             `json:"b_choice"`
	CChoice       string                             `json:"c_choice"`
	DChoice       string                             `json:"d_choice"`
	CorrectChoice domain.GrammarLessonExerciseChoice `json:"correct_choice"`
	Explanation   string                             `json:"explanation"`
}

type GrammarLessonGetStatusOKResponse struct {
	Title    string                 `json:"title"`
	Meaning  string                 `json:"meaning"`
	Sections []GrammarLessonSection `json:"sections"`
	Exercise string                 `json:"exercise"`
}

type GrammarLessonExercisesGetStatusOKResponse struct {
	Exercises []GrammarLessonExercise `json:"exercises"`
}

type GrammarLessonExercisesPostStatusOKResponse struct {
	OverallScore float32 `json:"overall_score"`
}
