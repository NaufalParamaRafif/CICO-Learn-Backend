package domain

type MemorizeVocabularyDraft struct {
	ID                    uint
	Name                  string
	Image                 string
	TotalVocabulary       uint16
	HaveLearnedVocabulary uint16
	MustLearnedVocabulary uint16
	UserID                uint
}
