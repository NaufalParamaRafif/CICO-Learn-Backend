package domain

type GrammarSection struct {
	ID              uint
	Title           string
	SectionOrder    uint8
	GrammarContents []GrammarContent
	GrammarLessonID uint
	Exerccises      []GrammarExercise
}
