package domain

import (
	"context"
	"time"
)

type UserTypeSentenceContainVocabularyTerm struct {
	ID               uint
	FirstSentence    string
	SecondSentence   string
	ThirdSentence    string
	Score            float32
	Feedback         string
	VocabularyTermID uint
	UserID           uint
	CreatedAt        time.Time
}

type UserTypeSentenceContainVocabularyTermRepository interface {
	GetByID(ctx context.Context, userTypeSentenceContainVocabularyTermID uint) (UserTypeSentenceContainVocabularyTerm, error)
	GetByUserID(ctx context.Context, userID uint) ([]UserTypeSentenceContainVocabularyTerm, error)
	GetByVocabularyTermID(ctx context.Context, vocabularyTermID uint) ([]UserTypeSentenceContainVocabularyTerm, error)
	GetByUserIDAndVocabularyTermID(ctx context.Context, userID uint, vocabularyTermID uint) (UserTypeSentenceContainVocabularyTerm, error)

	Insert(ctx context.Context, userTypeSentenceContainVocabularyTerm *UserTypeSentenceContainVocabularyTerm) error
	Update(ctx context.Context, userTypeSentenceContainVocabularyTerm *UserTypeSentenceContainVocabularyTerm) error
	Delete(ctx context.Context, userTypeSentenceContainVocabularyTermID uint) error
}
