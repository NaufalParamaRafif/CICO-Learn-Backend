package domain

import (
	"context"
	"time"
)

type MemorizeVocabularyDraftSpeakSentenceContainVocabularyTerm struct {
	ID                        uint      `db:"id"`
	Accuracy                  float32   `db:"accuracy"`
	Feedback                  string    `db:"feedback"`
	CreatedAt                 time.Time `db:"created_at"`
	UpdatedAt                 time.Time `db:"updated_at"`
	VocabularyTermID          uint      `db:"vocabulary_term_id"`
	MemorizeVocabularyDraftID uint      `db:"memorize_vocabulary_draft_id"`
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
