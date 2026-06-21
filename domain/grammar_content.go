package domain

import "context"

type GrammarContent struct {
	ID               uint                         `db:"id"`
	Title            string                       `db:"title"`
	Order            uint8                        `db:"grammar_content_order"`
	Content          string                       `db:"content"`
	Type             GrammarLessonExplanationType `db:"type"`
	GrammarSectionID uint                         `db:"grammar_section_id"`
}

type GrammarContentRepository interface {
	GetByID(ctx context.Context, grammarContentID uint) (GrammarContent, error)
	Insert(ctx context.Context, grammarContent *GrammarContent) error
	Update(ctx context.Context, grammarContent *GrammarContent) error
	Delete(ctx context.Context, grammarContentID uint) error
}
