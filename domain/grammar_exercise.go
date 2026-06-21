package domain

import "context"

type GrammarExercise struct {
	ID                   uint                        `db:"id"`
	GrammarExerciseOrder uint8                       `db:"grammar_exercise_order"`
	Question             string                      `db:"question"`
	AChoice              string                      `db:"a_choice"`
	BChoice              string                      `db:"b_choice"`
	CChoice              string                      `db:"c_choice"`
	DChoice              string                      `db:"d_choice"`
	CorrectAnswer        GrammarLessonExerciseChoice `db:"correct_answer"`
	GrammarSectionID     uint                        `db:"grammar_section_id"`
}

type GrammarExerciseRepository interface {
	GetByID(ctx context.Context, grammarExerciseID uint) (GrammarExercise, error)
	Insert(ctx context.Context, grammarExercise *GrammarExercise) error
	Update(ctx context.Context, grammarExercise *GrammarExercise) error
	Delete(ctx context.Context, grammarExerciseID uint) error
}
