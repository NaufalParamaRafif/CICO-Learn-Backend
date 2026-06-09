package dto

import "github.com/NaufalParamaRafif/CICO-Learn-Backend/domain"

type Lesson interface {
	GetLessonType() domain.LessonType
}

func (v VocabularyComponent) GetLessonType() domain.LessonType {
	return v.LessonType
}
func (g GrammarComponent) GetLessonType() domain.LessonType {
	return g.LessonType
}
func (v DialogComponent) GetLessonType() domain.LessonType {
	return v.LessonType
}

type VocabularyComponent struct {
	LessonType      domain.LessonType `json:"lesson_type"`
	Image           string            `json:"image"`
	Title           string            `json:"title"`
	TotalVocabulary uint16            `json:"total_vocabulary"`
	Type            string            `json:"type"`
	Progress        float32           `json:"progress"`
}

type GrammarComponent struct {
	LessonType    domain.LessonType `json:"lesson_type"`
	Image         string            `json:"image"`
	Title         string            `json:"title"`
	TotalSection  uint8             `json:"total_section"`
	TotalExercise uint16            `json:"total_exercise"`
	Progress      float32           `json:"progress"`
}

type DialogComponent struct {
	LessonType  domain.LessonType `json:"lesson_type"`
	Image       string            `json:"image"`
	Title       string            `json:"title"`
	TotalDialog uint16            `json:"total_dialog"`
	Progress    float32           `json:"progress"`
}
