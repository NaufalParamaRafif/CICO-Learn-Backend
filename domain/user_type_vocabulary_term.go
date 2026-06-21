package domain

import (
	"context"
	"time"
)

type UserTypeVocabularyTerm struct {
	ID               uint      `db:"id"`
	IsCorrect        bool      `db:"is_correct"`
	UserID           uint      `db:"user_id"`
	VocabularyTermID uint      `db:"vocabulary_term_id"`
	CreatedAt        time.Time `db:"created_at"`
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
