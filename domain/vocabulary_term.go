package domain

import "context"

type VocabularyTerm struct {
	Image    string             `db:"image"`
	Term     string             `db:"term"`
	Meaning  string             `db:"meaning"`
	Synonyms string             `db:"synonyms"`
	Sentence string             `db:"sentence"`
	TermType VocabularyTermType `db:"term_type"`
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
