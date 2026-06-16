package domain

import "context"

type UserGrammarExercise struct {
	ID                uint
	UserChoice        string
	IsCorrect         bool
	UserID            uint
	GrammarExerciseID uint
}

type UserGrammarExerciseRepository interface {
	GetByID(ctx context.Context, userGrammarExerciseID uint) (UserGrammarExercise, error)
	GetByUserID(ctx context.Context, userID uint) ([]UserGrammarExercise, error)
	GetByGrammarExerciseID(ctx context.Context, grammarExerciseID uint) ([]UserGrammarExercise, error)
	GetByUserIDAndGrammarExerciseID(ctx context.Context, userID uint, grammarExerciseID uint) (UserGrammarExercise, error)

	Insert(ctx context.Context, userGrammarExercise *UserGrammarExercise) error
	Update(ctx context.Context, userGrammarExercise *UserGrammarExercise) error
	Delete(ctx context.Context, userGrammarExercise uint) error
}
