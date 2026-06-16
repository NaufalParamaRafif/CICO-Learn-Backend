package domain

import (
	"context"
	"time"
)

type UserTypeVocabularyTerm struct {
	ID               uint
	IsCorrect        bool
	UserID           uint
	VocabularyTermID uint
	CreatedAt        time.Time
}

type UserTypeVocabularyTermRepository interface {
	GetByID(ctx context.Context, userTypeVocabularyTermID uint) (UserTypeVocabularyTerm, error)
	GetByUserID(ctx context.Context, userID uint) ([]UserTypeVocabularyTerm, error)
	GetByVocabularyTermID(ctx context.Context, vocabularyTermID uint) ([]UserTypeVocabularyTerm, error)
	GetByUserIDAndVocabularyTermID(ctx context.Context, userID uint, vocabularyTermID uint) (UserTypeVocabularyTerm, error)

	Insert(ctx context.Context, userTypeVocabularyTerm *UserTypeVocabularyTerm) error
	Update(ctx context.Context, userTypeVocabularyTerm *UserTypeVocabularyTerm) error
	Delete(ctx context.Context, userTypeVocabularyTermID uint) error
}
