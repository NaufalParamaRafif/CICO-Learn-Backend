package domain

import "context"

type GrammarExercise struct {
	ID               uint
	Question         string
	AChoice          string
	BChoice          string
	CChoice          string
	DChoice          string
	CorrectAnswer    GrammarLessonExerciseChoice
	GrammarSectionID uint
}

type GrammarExerciseRepository interface {
	GetByID(ctx context.Context, grammarExerciseID uint) (GrammarExercise, error)
	Insert(ctx context.Context, grammarExercise *GrammarExercise) error
	Update(ctx context.Context, grammarExercise *GrammarExercise) error
	Delete(ctx context.Context, grammarExerciseID uint) error
}
