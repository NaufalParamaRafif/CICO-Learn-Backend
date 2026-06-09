package dto

type MemorizeDraft struct {
	Image                 string `json:"image"`
	Title                 string `json:"title"`
	VocabularyTotal       uint16 `json:"vocabulary_total"`
	VocabularyHaveLearned uint16 `json:"vocabulary_have_learned"`
	VocabularyMustLearned uint16 `json:"vocabulary_must_learned"`
}

type MemorizeHistory struct {
	Image                 string `json:"image"`
	Title                 string `json:"title"`
	LastOpened            string `json:"last_opened"` // FORMAT: TIME WITH TIMEZONE
	VocabularyTotal       uint16 `json:"vocabulary_total"`
	VocabularyHaveLearned uint16 `json:"vocabulary_have_learned"`
	VocabularyMustLearned uint16 `json:"vocabulary_must_learned"`
}

type MemorizeGetStatusOKResponse struct {
	TodayVocabulary MemorizeDraft   `json:"today_vocabulary"`
	Drafts          []MemorizeDraft `json:"drafts"`
}

type MemorizeHistoriesGetStatusOKResponse struct {
	Histories []MemorizeHistory `json:"histories"`
}
