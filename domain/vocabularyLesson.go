package domain

type VocabularyLesson struct {
	ID              uint
	Topic           string
	Description     string
	Type            LessonType
	TotalVocabulary uint16
	Words           []Word
	Expression      []Expression
}
