package response

type VocabularyTerm struct {
	Word    string `json:"word"`
	Meaning string `json:"meaning"`
	IsSaved bool   `json:"is_saved"`
}

type VocabularyTermDetail struct {
	Word     string `json:"word"`
	WordInUK string `json:"word_in_uk"`
	WordInUS string `json:"word_in_us"`
	Meaning  string `json:"meaning"`
	Synonyms string `json:"synonyms"`
	// TODO: Add sound
}

type VocabularyExample struct {
	Example string `json:"example"`
	Meaning string `json:"meaning"`
}

type VocabularyLessonSentences struct {
	TypedSentence     string `json:"typed_sentence"`
	CorrectedSentence string `json:"corrected_sentence"`
	Explanation       string `json:"explanation"`
}

type VocabularyLessonTopicGetStatusOKResponse struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Vocabulary  []VocabularyTerm `json:"vocabulary"`
}

type VocabularyLessonExamplesGetStatusOKResponse struct {
	Vocabulary VocabularyTermDetail `json:"vocabulary"`
	Examples   []VocabularyExample  `json:"examples"`
}

type VocabularyLessonOverallScoreGetStatusOKResponse struct {
	Topic          string  `json:"topic"`
	VocabularyTerm string  `json:"vocabulary_term"`
	OverallScore   float32 `json:"overall_score"`
}

type VocabularyLessonTypeWordPostStatusOKResponse struct {
	CorrectWord string `json:"correct_word"`
	TypedWord   string `json:"typed_word"`
	IsCorrect   bool   `json:"is_correct"`
}

type VocabularyLessonWriteThreeSentencesPostStatusOKResponse struct {
	Sentences []VocabularyLessonSentences `json:"sentences"`
}
