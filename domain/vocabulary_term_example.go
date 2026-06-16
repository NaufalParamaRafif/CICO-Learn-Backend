package domain

import "context"

type VocabularyTermExample struct {
	ID               uint
	Example          string
	Meaning          string
	VocabularyTermID uint
}

type WordExampleRepository interface {
	GetByID(ctx context.Context, vocabularyTermExampleID uint) (VocabularyTermExample, error)
	Insert(ctx context.Context, vocabularyTerm *VocabularyTerm) error
	Update(ctx context.Context, vocabularyTerm *VocabularyTerm) error
	Delete(ctx context.Context, vocabularyTermExampleID uint) error
}
