package domain

import (
	"context"
	"time"
)

type MemorizeVocabularyDraftTypeTwoSentencesContainVocabularyTerm struct {
	ID                        uint
	FirstSentence             string
	SecondSentece             string
	Score                     float32
	Feedback                  string
	VocabularyTermID          uint
	MemorizeVocabularyDraftID uint
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

type MemorizeVocabularyDraftTypeTwoSentencesContainVocabularyTermRepository interface {
	GetByID(ctx context.Context, memorizeVocabularyDraftTypeTwoSentencesContainVocabularyTermID uint) (MemorizeVocabularyDraftTypeTwoSentencesContainVocabularyTerm, error)
	GetByVocabularyTermID(ctx context.Context, vocabularyTermID uint) ([]MemorizeVocabularyDraftTypeTwoSentencesContainVocabularyTerm, error)
	GetByMemorizeVocabularyDraftID(ctx context.Context, memorizeVocabularyDraftID uint) ([]MemorizeVocabularyDraftTypeTwoSentencesContainVocabularyTerm, error)
	GetByVocabularyTermIDAndMemorizeVocabularyDraftID(ctx context.Context, vocabularyTermID uint, memorizeVocabularyDraftID uint) (MemorizeVocabularyDraftTypeTwoSentencesContainVocabularyTerm, error)

	Insert(ctx context.Context, memorizeVocabularyDraftTypeTwoSentencesContainVocabularyTerm *MemorizeVocabularyDraftTypeTwoSentencesContainVocabularyTerm) error
	Update(ctx context.Context, memorizeVocabularyDraftTypeTwoSentencesContainVocabularyTerm *MemorizeVocabularyDraftTypeTwoSentencesContainVocabularyTerm) error
	Delete(ctx context.Context, memorizeVocabularyDraftTypeTwoSentencesContainVocabularyTermID uint) error
}
