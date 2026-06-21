package domain

import (
	"context"
	"time"
)

type UserTypeSentenceContainVocabularyTerm struct {
	ID               uint      `db:"id"`
	FirstSentence    string    `db:"first_sentence"`
	SecondSentence   string    `db:"second_sentence"`
	ThirdSentence    string    `db:"third_sentence"`
	Score            float32   `db:"score"`
	Feedback         string    `db:"feedback"`
	VocabularyTermID uint      `db:"vocabulary_term_id"`
	UserID           uint      `db:"user_id"`
	CreatedAt        time.Time `db:"created_at"`
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
