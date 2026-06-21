package domain

import (
	"context"
	"time"
)

type MemorizeVocabularyDraftTypeVocabularyTerm struct {
	ID                        uint      `db:"id"`
	IsCorrect                 bool      `db:"is_correct"`
	CreatedAt                 time.Time `db:"created_at"`
	UpdatedAt                 time.Time `db:"updated_at"`
	VocabularyTermID          uint      `db:"vocabulary_term_id"`
	MemorizeVocabularyDraftID uint      `db:"memorize_vocabulary_draft_id"`
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
