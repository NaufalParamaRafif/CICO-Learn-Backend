package domain

type Expression struct {
	ID       uint
	Phrase   string
	Meaning  string
	Sentence string
	Type     ExpressionType
	// VocabularyID uint // TODO: FIX THIS
	Examples []ExpressionExample
}
