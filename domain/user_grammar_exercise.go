package domain

import "context"

type UserGrammarExercise struct {
	ID                uint   `db:"id"`
	UserChoice        string `db:"user_choice"`
	IsCorrect         bool   `db:"is_correct"`
	UserID            uint   `db:"user_id"`
	GrammarExerciseID uint   `db:"grammar_exercise_id"`
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
