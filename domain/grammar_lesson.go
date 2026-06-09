package domain

type GrammarLesson struct {
	ID              uint
	Topic           string
	Description     string
	GrammarSections []GrammarSection // TODO: Fix Migrations
	Exercises       []GrammarExercise
}
