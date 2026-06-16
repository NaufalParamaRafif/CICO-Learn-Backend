package domain

import "context"

type VocabularyLesson struct {
	ID          uint
	Topic       string
	Description string
	Type        LessonType
	// TotalVocabulary uint16
	// Term []VocabularyTerm
}

// FIX DUAL TYPE (JUST TERM)

type VocabularyLessonRepository interface {
	GetByID(ctx context.Context, vocabularyLessonID uint) (VocabularyLesson, error)
	Insert(ctx context.Context, vocabularyLesson *VocabularyLesson) error
	Update(ctx context.Context, vocabularyLesson *VocabularyLesson) error
	Delete(ctx context.Context, vocabularyLessonID uint) error

	GetVocabularyTerms(ctx context.Context, vocabularyLessonID *VocabularyTerm) ([]VocabularyTerm, error)
	AddVocabularyTerm(ctx context.Context, vocabularyTermID uint) error
	RemoveVocabularyTerm(ctx context.Context, vocabularyTermID uint) error
}
