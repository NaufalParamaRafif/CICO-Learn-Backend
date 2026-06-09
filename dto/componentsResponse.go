package dto

type Lesson interface {
	GetLessonType() LessonType
}

func (v VocabularyComponent) GetLessonType() LessonType {
	return v.LessonType
}
func (g GrammarComponent) GetLessonType() LessonType {
	return g.LessonType
}
func (v DialogComponent) GetLessonType() LessonType {
	return v.LessonType
}

type VocabularyComponent struct {
	LessonType      LessonType `json:"lesson_type"`
	Image           string     `json:"image"`
	Title           string     `json:"title"`
	TotalVocabulary uint16     `json:"total_vocabulary"`
	Type            string     `json:"type"`
	Progress        float32    `json:"progress"`
}

type GrammarComponent struct {
	LessonType    LessonType `json:"lesson_type"`
	Image         string     `json:"image"`
	Title         string     `json:"title"`
	TotalSection  uint8      `json:"total_section"`
	TotalExercise uint16     `json:"total_exercise"`
	Progress      float32    `json:"progress"`
}

type DialogComponent struct {
	LessonType  LessonType `json:"lesson_type"`
	Image       string     `json:"image"`
	Title       string     `json:"title"`
	TotalDialog uint16     `json:"total_dialog"`
	Progress    float32    `json:"progress"`
}
