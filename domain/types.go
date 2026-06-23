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
type VocabularyLessonType string
type UserImageType string
type VocabularyTermType string

type FileInput struct {
	Filename    string
	ContentType string
	Size        float32 // KB
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
	DailyReminder   []TimeOfDay
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

	// Gender
	Men     Gender = "men"
	Women   Gender = "women"
	Unknown Gender = "unknown"

	// EnglishLevel
	Fondation         EnglishLevel = "Pre-A1(foundation)"
	Elementary        EnglishLevel = "A1(Elementary)"
	PreIntermediate   EnglishLevel = "A2(Pre-Intermediate)"
	Intermediate      EnglishLevel = "B1(Intermediate)"
	UpperIntermediate EnglishLevel = "B2(Upper Intermediate)"
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
	FiveMinutes      TimeTarget = "5 minutes"
	TenMinutes       TimeTarget = "10 minutes"
	FifteenMinutes   TimeTarget = "15 minutes"
	ThirtyMinutes    TimeTarget = "30 minutes"
	FortyFiveMinutes TimeTarget = "45 minutes"
	OneHour          TimeTarget = "1 hour"
	TwoHours         TimeTarget = "2 hours"

	// TrainWithAIMode
	FreeTalk  TrainWithAIMode = "free_talk"
	RolePlay  TrainWithAIMode = "role_play"
	Challange TrainWithAIMode = "challenge"

	// GrammarLessonExplanationType
	Header      GrammarLessonExplanationType = "header"
	Explanation GrammarLessonExplanationType = "text"

	// GrammarLessonExerciseChoice
	AChoice GrammarLessonExerciseChoice = "a"
	BChoice GrammarLessonExerciseChoice = "b"
	CChoice GrammarLessonExerciseChoice = "c"
	DChoice GrammarLessonExerciseChoice = "d"

	// VocabularyTermType
	Word         VocabularyTermType = "word"
	PhrasalVerbs VocabularyTermType = "phrasal_verbs"
	Collocations VocabularyTermType = "collocations"
	Phrases      VocabularyTermType = "phrases"
	Idoms        VocabularyTermType = "idoms"
	Sayins       VocabularyTermType = "sayins"
	Proverbs     VocabularyTermType = "proverbs"

	// UserImageType
	UserImageProfile    UserImageType = "profile"
	UserImageBackground UserImageType = "background"
)

func (eL EnglishLevel) IsValid() bool {
	switch eL {
	case Fondation,
		Elementary,
		PreIntermediate,
		Intermediate,
		UpperIntermediate,
		Advanced,
		Proficiency:
		return true
	default:
		return false
	}
}

func (g Gender) IsValid() bool {
	switch g {
	case Men, Women, Unknown:
		return true
	default:
		return false
	}
}

func (l Language) IsValid() bool {
	switch l {
	case Indonesia, English:
		return true
	default:
		return false
	}
}

func (trgt Target) IsValid() bool {
	switch trgt {
	case JobInterview,
		LearningForFun,
		StudyAbroad,
		FluentCommunication,
		UnderstandUniversityLecture,
		Work:
		return true
	default:
		return false
	}
}

func (tmTrgt TimeTarget) IsValid() bool {
	switch tmTrgt {
	case FiveMinutes,
		TenMinutes,
		FifteenMinutes,
		ThirtyMinutes,
		FortyFiveMinutes,
		OneHour,
		TwoHours:
		return true
	default:
		return false
	}
}
