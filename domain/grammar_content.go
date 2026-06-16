package domain

import "context"

type GrammarContent struct {
	ID               uint
	Title            string
	Order            uint8
	Content          string
	Type             GrammarLessonExplanationType
	GrammarSectionID uint
}

type GrammarContentRepository interface {
	GetByID(ctx context.Context, grammarContentID uint) (GrammarContent, error)
	Insert(ctx context.Context, grammarContent *GrammarContent) error
	Update(ctx context.Context, grammarContent *GrammarContent) error
	Delete(ctx context.Context, grammarContentID uint) error
}
