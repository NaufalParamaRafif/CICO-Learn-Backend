package domain

import (
	"context"
	"time"
)

type UserSpeakSentenceContainVocabularyTerm struct {
	ID               uint
	Accuracy         float32
	Feedback         string
	VocabularyTermID uint
	UserID           uint
	CreatedAt        time.Time
}

type UserSpeakSentenceContainVocabularyTermRepository interface {
	GetByID(ctx context.Context, userSpeakSenteceContainVocabularyTermID uint) (UserSpeakSentenceContainVocabularyTerm, error)
	GetByUserID(ctx context.Context, userID uint) ([]UserSpeakSentenceContainVocabularyTerm, error)
	GetByVocabularyTermID(ctx context.Context, vocabularyTermID uint) ([]UserSpeakSentenceContainVocabularyTerm, error)
	GetByUserIDAndVocabularyTermID(ctx context.Context, userID uint, vocabularyTermID uint) (UserSpeakSentenceContainVocabularyTerm, error)

	Insert(ctx context.Context, userSpeakSenteceContainVocabularyTerm *UserSpeakSentenceContainVocabularyTerm) error
	Update(ctx context.Context, userSpeakSenteceContainVocabularyTerm *UserSpeakSentenceContainVocabularyTerm) error
	Delete(ctx context.Context, userSpeakSenteceContainVocabularyTermID uint) error
}
