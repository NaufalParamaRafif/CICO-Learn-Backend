package domain

import "context"

type MemorizeVocabularyDraft struct {
	ID    uint
	Name  string
	Image string
	// TotalVocabulary       uint16
	// HaveLearnedVocabulary uint16
	// MustLearnedVocabulary uint16
}

type MemorizeVocabularyDraftRepository interface {
	GetByID(ctx context.Context, memorizeVocabularyDraftID uint) (MemorizeVocabularyDraft, error)
	Insert(ctx context.Context, memorizeVocabularyDraft *MemorizeVocabularyDraft) error
	Update(ctx context.Context, memorizeVocabularyDraft *MemorizeVocabularyDraft) error
	Delete(ctx context.Context, memorizeVocabularyDraftID uint) error

	GetVocabularyTerms(ctx context.Context, vocabularyTermID uint) ([]MemorizeVocabularyDraft, error)
	AddVocabularyTerm(ctx context.Context, vocabularyTermID uint) error
	DeleteVocabularyTerm(ctx context.Context, vocabularyTermID uint) error

	GetMemorizeVocabularyDraftTypeVocabularyTerms(ctx context.Context, memorizeVocabularyDraftID uint) ([]MemorizeVocabularyDraftTypeVocabularyTerm, error)
	GetMemorizeVocabularyDraftTypeTwoSentencesContainVocabularyTerm(ctx context.Context, memorizeVocabularyDraftID uint) ([]MemorizeVocabularyDraftTypeTwoSentencesContainVocabularyTerm, error)
	GetMemorizeVocabularyDraftSpeakSentenceContainVocabularyTerm(ctx context.Context, memorizeVocabularyDraftID uint) ([]MemorizeVocabularyDraftSpeakSentenceContainVocabularyTerm, error)
}
