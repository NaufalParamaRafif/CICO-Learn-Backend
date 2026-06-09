package domain

type GrammarContent struct {
	ID               uint
	Title            string
	Order            uint8
	Content          string
	Type             GrammarLessonExplanationType
	GrammarSectionID uint
}
