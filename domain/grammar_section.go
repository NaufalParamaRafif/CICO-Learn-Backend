package domain

import "context"

type GrammarSection struct {
	ID              uint
	Title           string
	SectionOrder    uint8
	GrammarLessonID uint
	GrammarContents []GrammarContent
	Exercises       []GrammarExercise
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
