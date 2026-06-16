package domain

import (
	"context"
	"time"
)

type MemorizeVocabularyDraftSpeakSentenceContainVocabularyTerm struct {
	ID                        uint
	Accuracy                  float32
	Feedback                  string
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	VocabularyTermID          uint
	MemorizeVocabularyDraftID uint
}

type MemorizeVocabularyDraftSpeakSentenceContainVocabularyTermRepository interface {
	GetByID(ctx context.Context, memorizeVocabularyDraftSpeakSentenceContainVocabularyTermID uint) (MemorizeVocabularyDraftSpeakSentenceContainVocabularyTerm, error)
	GetByVocabularyTermID(ctx context.Context, vocabularyTermID uint) ([]MemorizeVocabularyDraftSpeakSentenceContainVocabularyTerm, error)
	GetByMemorizeVocabularyDraftID(ctx context.Context, memorizeVocabularyDraftID uint) ([]MemorizeVocabularyDraftSpeakSentenceContainVocabularyTerm, error)
	GetByVocabularyTermIDAndMemorizeVocabularyDraftID(ctx context.Context, vocabularyTermID uint, memorizeVocabularyDraftID uint) (MemorizeVocabularyDraftSpeakSentenceContainVocabularyTerm, error)

	Insert(ctx context.Context, memorizeVocabularyDraftSpeakSentenceContainVocabularyTerm *MemorizeVocabularyDraftSpeakSentenceContainVocabularyTerm) error
	Update(ctx context.Context, memorizeVocabularyDraftSpeakSentenceContainVocabularyTerm *MemorizeVocabularyDraftSpeakSentenceContainVocabularyTerm) error
	Delete(ctx context.Context, memorizeVocabularyDraftSpeakSentenceContainVocabularyTermID uint) error
}
