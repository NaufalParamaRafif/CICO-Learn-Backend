package domain

import "context"

type GrammarSection struct {
	ID              uint              `db:"id"`
	Title           string            `db:"title"`
	SectionOrder    uint8             `db:"section_order"`
	GrammarLessonID uint              `db:"grammar_lesson_id"`
	GrammarContents []GrammarContent  `db:"grammar_contents"`
	Exercises       []GrammarExercise `db:"exercises"`
}

type GrammarSectionRepository interface {
	GetByID(ctx context.Context, grammarSectionID uint) (GrammarSection, error)
	Insert(ctx context.Context, grammarSection *GrammarSection) error
	Update(ctx context.Context, grammarSection *GrammarSection) error
	Delete(ctx context.Context, grammarSectionID uint) error

	GetUsers(ctx context.Context, grammarSectionID uint) ([]User, error)
	AddUser(ctx context.Context, userID uint) error
	RemoveUser(ctx context.Context, userID uint) error

	GetExercises(ctx context.Context, grammarSectionID uint) ([]GrammarExercise, error)
	GetContents(ctx context.Context, grammarSectionID uint) ([]GrammarContent, error)
}
