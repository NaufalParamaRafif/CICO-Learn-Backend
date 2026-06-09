package dto

type VocabularyLessonSpeakGetStatusOKResponse struct {
	Sentence        string `json:"sentence"`
	HighlightedWord string `json:"highlighted_word"`
}

type VocabularyLessonSpeakPostStatusOKResponse struct {
	Sentence        string  `json:"sentence"`
	HighlightedWord string  `json:"highlighted_word"`
	Accuracy        float32 `json:"accuracy"`
	Meaning         string  `json:"meaning"`
}
