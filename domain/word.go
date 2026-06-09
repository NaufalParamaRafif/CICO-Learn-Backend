package domain

type Word struct {
	Image    string
	Word     string
	Meaning  string
	Synonyms string
	Sentence string
	// VocabularyID uint // TODO: FIX THIS
	Examples []WordExample
}
