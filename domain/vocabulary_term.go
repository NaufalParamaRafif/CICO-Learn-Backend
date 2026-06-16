package domain

import "context"

type VocabularyTerm struct {
	Image          string
	VocabularyTerm string // TODO: FIX MIGRATIONS
	Meaning        string
	Synonyms       string
	Sentence       string
	// Examples []VocabularyTermExample
	// TODO: ADD SOUND
}

type VocabularyTermRepository interface {
	GetByID(ctx context.Context, vocabularyTermID VocabularyTerm) (VocabularyTerm, error)
	Insert(ctx context.Context, vocabularyTerm *VocabularyTerm) error
	Update(ctx context.Context, vocabularyTerm *VocabularyTerm) error
	Delete(ctx context.Context, vocabularyTermID uint) error

	GetExamples(ctx context.Context, vocabularyTermID uint) ([]VocabularyTermExample, error)
}
