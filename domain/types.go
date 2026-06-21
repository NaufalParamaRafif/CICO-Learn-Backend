package domain

import (
	"io"
)

type Day string
type LessonType string
type EnglishLevel string
type Language string
type Target string
type Gender string
type TimeTarget string
type TrainWithAIMode string
type GrammarLessonExplanationType string
type GrammarLessonExerciseChoice string
type ExpressionType string
type VocabularyLessonType string

type FileInput struct {
	Filename    string
	ContentType string // JPEG, JPG, IMG
	Size        int64  // KB
	Data        io.Reader
}

type LoginResult struct {
	AccessToken  string
	RefreshToken string
}

type LoginInput struct {
	Identifier string
	Password   string
}

type RegisterInput struct {
	FirstName *string
	LastName  *string
	Username  string
	Email     string
	Password  string
}

type UpdateUserInput struct {
	ID              uint
	FirstName       *string
	LastName        *string
	Address         *string
	Gender          *Gender
	ProfilePicture  *string
	BackgroundImage *string
	Level           *EnglishLevel
	Language        *Language
	Target          *Target
	DailyReminder   *[]TimeOfDay
	TimeTarget      *TimeTarget
}

const (
	// Day
	Sunday    Day = "sunday"
	Monday    Day = "monday"
	Tuesday   Day = "tuesday"
	Wednesday Day = "wednesday"
	Thursday  Day = "thursday"
	Friday    Day = "friday"
	Saturday  Day = "saturday"

	// LessonType
	Vocabulary LessonType = "vocabulary"
	Grammar    LessonType = "grammar"
	Dialog     LessonType = "dialog"

	// EnglishLevel
	Fondation         EnglishLevel = "Pre-A1(foundation)"
	Elementary        EnglishLevel = "A1(Elementary)"
	PreIntermediate   EnglishLevel = "A2(Pre-Intermediate)"
	Intermediate      EnglishLevel = "B1(Intermediate)"
	UpperIntermediate EnglishLevel = "B2(Upper (Intermediate)"
	Advanced          EnglishLevel = "C1(Advanced)"
	Proficiency       EnglishLevel = "C2(Proficiency)"

	// Language
	English   Language = "english"
	Indonesia Language = "indonesia"

	// Target
	LearningForFun              Target = "Learning for fun"
	JobInterview                Target = "Job interview"
	StudyAbroad                 Target = "Study abroad"
	FluentCommunication         Target = "Fluent communication"
	UnderstandUniversityLecture Target = "Understand university lecture"
	Work                        Target = "Work"

	// TimeTarget
	FiveMinutes      TimeTarget = "5 Minutes"
	TenMinutes       TimeTarget = "10 Minutes"
	FifteenMinutes   TimeTarget = "15 Minutes"
	ThirtyMinutes    TimeTarget = "30 Minutes"
	FortyFiveMinutes TimeTarget = "45 Minutes"
	OneHour          TimeTarget = "1 Hour"
	TwoHours         TimeTarget = "2 Hours"

	// TrainWithAIMode
	FreeTalk  TrainWithAIMode = "free_talk"
	RolePlay  TrainWithAIMode = "role_play"
	Challange TrainWithAIMode = "challenge"

	// GrammarLessonExplanationType
	Header      GrammarLessonExplanationType = "header"
	Explanation GrammarLessonExplanationType = "explanation"

	// GrammarLessonExerciseChoice
	AChoice GrammarLessonExerciseChoice = "a"
	BChoice GrammarLessonExerciseChoice = "b"
	CChoice GrammarLessonExerciseChoice = "c"
	DChoice GrammarLessonExerciseChoice = "d"

	// ExpressionType
	PhrasalVerbs ExpressionType = "phrasal_verbs"
	Collocations ExpressionType = "collocations"
	Phrases      ExpressionType = "phrases"
	Idoms        ExpressionType = "idoms"
	Sayins       ExpressionType = "sayins"
	Proverbs     ExpressionType = "proverbs"

	// VocabularyLessonType
	ExpressionVocabulary VocabularyLessonType = "expression"
	WordVocabulary       VocabularyLessonType = "word"
)
