package domain

type GrammarExercise struct {
	ID               uint
	Question         string
	AChoice          string
	BChoice          string
	CChoice          string
	DChoice          string
	CorrectAnswer    GrammarLessonExerciseChoice
	GrammarSectionID uint
	GrammarLessonID  uint
}
