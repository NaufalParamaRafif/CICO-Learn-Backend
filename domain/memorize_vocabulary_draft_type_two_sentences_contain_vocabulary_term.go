package domain

import (
	"context"
	"time"
)

type MemorizeVocabularyDraftTypeTwoSentencesContainVocabularyTerm struct {
	ID                        uint      `db:"id"`
	FirstSentence             string    `db:"first_sentence"`
	SecondSentece             string    `db:"second_sentece"`
	Score                     float32   `db:"score"`
	Feedback                  string    `db:"feedback"`
	VocabularyTermID          uint      `db:"vocabulary_term_id"`
	MemorizeVocabularyDraftID uint      `db:"memorize_vocabulary_draft_id"`
	CreatedAt                 time.Time `db:"created_at"`
	UpdatedAt                 time.Time `db:"updated_at"`
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
