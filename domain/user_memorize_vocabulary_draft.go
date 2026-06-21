package domain

import (
	"context"
	"time"
)

type UserMemorizeVocabularyDraft struct {
	ID                        uint      `db:"id"`
	UserID                    uint      `db:"user_id"`
	MemorizeVocabularyDraftID uint      `db:"memorize_vocabulary_draft_id"`
	CreatedAt                 time.Time `db:"created_at"`
}

type UserMemorizeVocabularyDraftID interface {
	GetByID(ctx context.Context, userMemorizeVocabularyDraftID uint) (UserMemorizeVocabularyDraft, error)
	GetByUserID(ctx context.Context, userID uint) ([]UserMemorizeVocabularyDraft, error)
	GetByMemorizeVocabularyDraftID(ctx context.Context, memorizeVocabularyDraftID uint) ([]UserMemorizeVocabularyDraft, error)
	GetByUserIDAndMemorizeVocabularyDraftID(ctx context.Context, userID uint, memorizeVocabularyDraftID uint) (UserMemorizeVocabularyDraft, error)

	Insert(ctx context.Context, userMemorizeVocabularyDraft *UserMemorizeVocabularyDraft) error
	Update(ctx context.Context, userMemorizeVocabularyDraft *UserMemorizeVocabularyDraft) error
	Delete(ctx context.Context, userMemorizeVocabularyDraftID uint) error

	GetMemorizeVocabularyDrafts(ctx context.Context, userMemorizeVocabularyDraftID uint) ([]MemorizeVocabularyDraft, error)
}
