package domain

import (
	"context"
	"time"
)

type MemorizeVocabularyDraftTypeVocabularyTerm struct {
	ID                        uint
	IsCorrect                 bool
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	VocabularyTermID          uint
	MemorizeVocabularyDraftID uint
}

type MemorizeVocabularyDraftTypeVocabularyTermsRepository interface {
	GetByID(ctx context.Context, memorizeVocabularyDraftTypeVocabularyTermID uint) (MemorizeVocabularyDraftTypeVocabularyTerm, error)
	GetByVocabularyTermID(ctx context.Context, vocabularyTermID uint) ([]MemorizeVocabularyDraftTypeVocabularyTerm, error)
	GetByMemorizeVocabularyDraftID(ctx context.Context, memorizeVocabularyDraftID uint) ([]MemorizeVocabularyDraftTypeVocabularyTerm, error)
	GetByVocabularyTermIDAndMemorizeVocabularyDraftID(ctx context.Context, vocabularyTermID uint, memorizeVocabularyDraftID uint) (MemorizeVocabularyDraftTypeVocabularyTerm, error)

	Insert(ctx context.Context, memorizeVocabularyDraftTypeVocabularyTerms *MemorizeVocabularyDraftTypeVocabularyTerm) error
	Update(ctx context.Context, memorizeVocabularyDraftTypeVocabularyTerms *MemorizeVocabularyDraftTypeVocabularyTerm) error
	Delete(ctx context.Context, memorizeVocabularyDraftTypeVocabularyTermsID uint) error
}
