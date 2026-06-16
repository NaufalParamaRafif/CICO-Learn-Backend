package domain

import "context"

type GrammarLesson struct {
	ID          uint
	Topic       string
	Description string
	// GrammarSections []GrammarSection
	// Exercises       []GrammarExercise
}

type GrammarLessonRepository interface {
	GetByID(ctx context.Context, grammarLessonID uint) (GrammarLesson, error)
	Insert(ctx context.Context, grammarLesson *GrammarLesson) error
	Update(ctx context.Context, grammarLesson *GrammarLesson) error
	Delete(ctx context.Context, grammarLessonID uint) error

	GetUsers(ctx context.Context, grammarLessonID uint) ([]User, error)
	AddUser(ctx context.Context, userID uint) error
	RemoveUser(ctx context.Context, userID uint) error

	GetSections(ctx context.Context, grammarLessonID uint) ([]GrammarSection, error)
}
