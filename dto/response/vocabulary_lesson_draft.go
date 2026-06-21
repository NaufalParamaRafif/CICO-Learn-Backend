package response

type VocabularyLessonDraft struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	IsSavedHere bool   `json:"is_saved_here"`
}

type VocabularyLessonDraftsGetStatusOKResponse struct {
	Drafts []VocabularyLessonDraft `json:"drafts"`
}

type VocabularyLessonDraftsPostStatusOKResponse struct {
	Word   string                  `json:"word"`
	Drafts []VocabularyLessonDraft `json:"drafts"`
}
