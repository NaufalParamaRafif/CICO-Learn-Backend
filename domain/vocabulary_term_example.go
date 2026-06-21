package domain

import "context"

type VocabularyTermExample struct {
	ID               uint   `db:"id"`
	Example          string `db:"example"`
	Meaning          string `db:"meaning"`
	VocabularyTermID uint   `db:"vocabulary_term_id"`
}

type WordExampleRepository interface {
	GetByID(ctx context.Context, vocabularyTermExampleID uint) (VocabularyTermExample, error)
	Insert(ctx context.Context, vocabularyTerm *VocabularyTerm) error
	Update(ctx context.Context, vocabularyTerm *VocabularyTerm) error
	Delete(ctx context.Context, vocabularyTermExampleID uint) error
}
